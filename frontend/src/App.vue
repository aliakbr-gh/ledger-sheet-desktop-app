<script setup lang="ts">
import { ref, onMounted, onUnmounted, reactive, watch, computed } from "vue";
import CashModal from "./CashModal.vue";
import { Toaster, toast } from 'vue-sonner'
import 'vue-sonner/style.css'
import pdfMake from "pdfmake/build/pdfmake";
import pdfFonts from "pdfmake/build/vfs_fonts";
import { SavePDF } from "../wailsjs/go/main/App";
import urduFontBase64 from "./assets/fonts/NotoNaskhArabic-Bold.base64.ts";

pdfMake.addVirtualFileSystem(pdfFonts.vfs as any);
pdfMake.addVirtualFileSystem({
    "NotoNaskhArabic-Bold.ttf": urduFontBase64,
});

pdfMake.fonts = {
    Roboto: {
        normal: 'Roboto-Regular.ttf',
        bold: 'Roboto-Regular.ttf',
        italics: 'Roboto-Regular.ttf',
        bolditalics: 'Roboto-Regular.ttf'
    },
    Urdu: {
        normal: "NotoNaskhArabic-Bold.ttf",
        bold: "NotoNaskhArabic-Bold.ttf",
        italics: "NotoNaskhArabic-Bold.ttf",
        bolditalics: "NotoNaskhArabic-Bold.ttf"
    }
};

const DEBUG = false;

type Sheet = {
    telenorOpeningBalance: number | null;
    jazzOpeningBalance: number | null;
    ufoneOpeningBalance: number | null;
    zongOpeningBalance: number | null;

    telenorNewBalance: number | null;
    jazzNewBalance: number | null;
    ufoneNewBalance: number | null;
    zongNewBalance: number | null;

    telenorReversalBalance: number | null;
    jazzReversalBalance: number | null;
    ufoneReversalBalance: number | null;
    zongReversalBalance: number | null;

    telenorClosingBalance: number | null;
    jazzClosingBalance: number | null;
    ufoneClosingBalance: number | null;
    zongClosingBalance: number | null;

    accountBalance265999891: number | null;
    accountBalance37300247: number | null;

    stampPaper: { name: string; amount: number | null }[];
    recovery: { name: string; amount: number | null }[];

    deposit265999891: number | null;
    deposit37300247: number | null;

    withdrawl265999891: number | null;
    withdrawl37300247: number | null;

    omni: { sending: string; receiving: string }[];
    easypaisa: { sending: string; receiving: string }[];
    jazzcash: { sending: string; receiving: string }[];
    epaccount: { sending: string; receiving: string }[];
    jcaccount: { sending: string; receiving: string }[];

    lastBalances: {
        omni: number | null;
        easypaisa: number | null;
        jazzcash: number | null;
        epaccount: number | null;
        jcaccount: number | null;
    };

    manualpurchasing: { name: string; amount: number | null }[];

    redBook: { name: string; amount: number | null }[];

    extra: number | null,

    previousCash: number | null,

    cash5000: number | null;
    cash1000: number | null;
    cash500: number | null;
    cash100: number | null;
    cash50: number | null;
    cash20: number | null;
    cash10: number | null;
    cash5: number | null;
};

const sheet = reactive<Sheet>({
    telenorOpeningBalance: null,
    jazzOpeningBalance: null,
    ufoneOpeningBalance: null,
    zongOpeningBalance: null,

    telenorNewBalance: null,
    jazzNewBalance: null,
    ufoneNewBalance: null,
    zongNewBalance: null,

    telenorReversalBalance: null,
    jazzReversalBalance: null,
    ufoneReversalBalance: null,
    zongReversalBalance: null,

    telenorClosingBalance: null,
    jazzClosingBalance: null,
    ufoneClosingBalance: null,
    zongClosingBalance: null,

    accountBalance265999891: null,
    accountBalance37300247: null,

    stampPaper: Array.from({ length: 13 }, () => ({
        name: "",
        amount: null as number | null,
    })),

    recovery: Array.from({ length: 12 }, () => ({
        name: "",
        amount: null as number | null,
    })),

    deposit265999891: null,
    deposit37300247: null,

    withdrawl265999891: null,
    withdrawl37300247: null,

    omni: Array.from({ length: 10 }, () => ({
        sending: "",
        receiving: "",
    })),

    easypaisa: Array.from({ length: 10 }, () => ({
        sending: "",
        receiving: "",
    })),

    jazzcash: Array.from({ length: 10 }, () => ({
        sending: "",
        receiving: "",
    })),

    epaccount: Array.from({ length: 10 }, () => ({
        sending: "",
        receiving: "",
    })),

    jcaccount: Array.from({ length: 10 }, () => ({
        sending: "",
        receiving: "",
    })),

    manualpurchasing: Array.from({ length: 9 }, () => ({
        name: "",
        amount: null,
    })),

    lastBalances: {
        omni: null,
        easypaisa: null,
        jazzcash: null,
        epaccount: null,
        jcaccount: null,
    },

    redBook: Array.from({ length: 8 }, () => ({
        name: "",
        amount: null as number | null,
    })),

    extra: null,

    previousCash: null,

    cash5000: null,
    cash1000: null,
    cash500: null,
    cash100: null,
    cash50: null,
    cash20: null,
    cash10: null,
    cash5: null,
});

const showCashModal = ref(false);
const showClearSheetModal = ref(false);

const savedSheet = localStorage.getItem("sheet");

if (savedSheet) {
    const parsed = JSON.parse(savedSheet);

    Object.assign(sheet, parsed);

    if (sheet.lastBalances) {
        sheet.lastBalances.omni = Number(sheet.lastBalances.omni) || null;
        sheet.lastBalances.easypaisa = Number(sheet.lastBalances.easypaisa) || null;
        sheet.lastBalances.jazzcash = Number(sheet.lastBalances.jazzcash) || null;
        sheet.lastBalances.epaccount = Number(sheet.lastBalances.epaccount) || null;
        sheet.lastBalances.jcaccount = Number(sheet.lastBalances.jcaccount) || null;
    }
}

watch(
    sheet,
    (newValue) => {
        localStorage.setItem(
            "sheet",
            JSON.stringify(newValue)
        );
    },
    {
        deep: true,
    }
);

const uploadBackup = () => {
    const input = document.createElement('input');
    input.type = 'file';
    input.accept = '.json';

    input.onchange = (e: Event) => {
        const file = (e.target as HTMLInputElement).files?.[0];
        if (!file) return;

        const reader = new FileReader();

        reader.onload = (event) => {
            try {
                const parsed = JSON.parse(event.target?.result as string);
                Object.assign(sheet, parsed);

                toast.success(`Backup "${file.name}" loaded successfully!`);
            } catch (err) {
                toast.error(`"${file.name}" is not a valid backup file`);
            }
        };

        reader.readAsText(file);
    };

    input.click();
};

const confirmClearSheet = () => {
    showClearSheetModal.value = false;

    sheet.previousCash = cashTotal.value;

    sheet.telenorOpeningBalance = sheet.telenorClosingBalance;
    sheet.jazzOpeningBalance = sheet.jazzClosingBalance;
    sheet.ufoneOpeningBalance = sheet.ufoneClosingBalance;
    sheet.zongOpeningBalance = sheet.zongClosingBalance;

    sheet.telenorNewBalance = null;
    sheet.jazzNewBalance = null;
    sheet.ufoneNewBalance = null;
    sheet.zongNewBalance = null;

    sheet.telenorReversalBalance = null;
    sheet.jazzReversalBalance = null;
    sheet.ufoneReversalBalance = null;
    sheet.zongReversalBalance = null;

    sheet.accountBalance265999891 =
        n(sheet.accountBalance265999891) +
        n(sheet.deposit265999891) -
        n(sheet.withdrawl265999891);

    sheet.accountBalance37300247 =
        n(sheet.accountBalance37300247) +
        n(sheet.deposit37300247) -
        n(sheet.withdrawl37300247);

    sheet.deposit265999891 = null;
    sheet.deposit37300247 = null;
    sheet.withdrawl265999891 = null;
    sheet.withdrawl37300247 = null;

    const updateLastBalance = (key: keyof typeof sheet.lastBalances, entries: any[]) => {
        const extracted = extractLastBalance(entries);
        if (extracted !== null) {
            sheet.lastBalances[key] = extracted;
        }
    };

    updateLastBalance('omni', sheet.omni);
    updateLastBalance('easypaisa', sheet.easypaisa);
    updateLastBalance('jazzcash', sheet.jazzcash);
    updateLastBalance('epaccount', sheet.epaccount);
    updateLastBalance('jcaccount', sheet.jcaccount);

    const clearEntries = (arr: any[]) => {
        arr.forEach((item) => {
            item.sending = "";
            item.receiving = "";
        });
    };

    clearEntries(sheet.omni);
    clearEntries(sheet.easypaisa);
    clearEntries(sheet.jazzcash);
    clearEntries(sheet.epaccount);
    clearEntries(sheet.jcaccount);

    sheet.manualpurchasing.forEach((item) => {
        item.name = "";
        item.amount = null;
    });

    sheet.stampPaper.forEach((item) => {
        item.name = "";
        item.amount = null;
    });

    sheet.recovery.forEach((item) => {
        item.name = "";
        item.amount = null;
    });

    sheet.redBook.forEach((item) => {
        item.name = "";
        item.amount = null;
    });

    sheet.extra = null;

    toast.success("Sheet cleared successfully");
};

// Helper Functions
const n = (v: number | null | undefined): number => {
    return Number(v) || 0;
};

const isPureNumber = (val: string | number | null | undefined) => {
    if (typeof val === "number") return true;

    if (typeof val === "string") {
        const cleaned = val.trim();

        return /^\d+(\.\d+)?$/.test(cleaned);
    }

    return false;
};

// Header
const dateTime = ref({
    date: "",
    islamicDate: "",
    pdfIslamicDate: "",
    time: "",
});

const updateDateTime = () => {
    const now = new Date();

    // English Date
    dateTime.value.date = now.toLocaleDateString("en-US", {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric",
    });

    // Islamic Urdu Date
    const islamicFormatter = new Intl.DateTimeFormat(
        "ur-PK-u-ca-islamic-umalqura",
        {
            weekday: "long",
            day: "numeric",
            month: "long",
            year: "numeric",
        }
    );

    const parts = islamicFormatter.formatToParts(now);

    const day = parts.find(p => p.type === "day")?.value || "";
    const month = parts.find(p => p.type === "month")?.value || "";
    const year = parts.find(p => p.type === "year")?.value || "";
    const weekday = parts.find(p => p.type === "weekday")?.value || "";

    dateTime.value.islamicDate = `${day} ${month} ${year} ہجری بروز ${weekday}`;
    dateTime.value.pdfIslamicDate = `${weekday} بروز ہجری ${year} ${month} ${day}`;

    // Time
    dateTime.value.time = now.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
    });
};
let interval: number;

onMounted(() => {
    const style = document.createElement("style");

    style.innerHTML = `
        @font-face {
            font-family: 'Urdu';
            src: url(data:font/truetype;charset=utf-8;base64,${urduFontBase64}) format('truetype');
            font-weight: normal;
            font-style: normal;
        }
    `;

    document.head.appendChild(style);

    updateDateTime();

    interval = window.setInterval(() => {
        updateDateTime();
    }, 1000);

    const preventWheelOnNumberInputs = (e: WheelEvent) => {
        const target = e.target as HTMLElement;
        if (
            target instanceof HTMLInputElement &&
            target.type === "number"
        ) {
            target.blur();
        }
    };

    document.addEventListener("wheel", preventWheelOnNumberInputs, {
        passive: true,
    });
});

onUnmounted(() => {
    clearInterval(interval);
});

// Sheet
const getELoadSell = (type: "telenor" | "jazz" | "ufone" | "zong") => {
    return (
        (sheet[`${type}OpeningBalance`] || 0) +
        (sheet[`${type}NewBalance`] || 0) +
        (sheet[`${type}ReversalBalance`] || 0) -
        (sheet[`${type}ClosingBalance`] || 0)
    );
};

const totalELoad = computed(() => {
    const telenor =
        (sheet.telenorOpeningBalance || 0) +
        (sheet.telenorNewBalance || 0) +
        (sheet.telenorReversalBalance || 0) -
        (sheet.telenorClosingBalance || 0);

    const jazz =
        (sheet.jazzOpeningBalance || 0) +
        (sheet.jazzNewBalance || 0) +
        (sheet.jazzReversalBalance || 0) -
        (sheet.jazzClosingBalance || 0);

    const ufone =
        (sheet.ufoneOpeningBalance || 0) +
        (sheet.ufoneNewBalance || 0) +
        (sheet.ufoneReversalBalance || 0) -
        (sheet.ufoneClosingBalance || 0);

    const zong =
        (sheet.zongOpeningBalance || 0) +
        (sheet.zongNewBalance || 0) +
        (sheet.zongReversalBalance || 0) -
        (sheet.zongClosingBalance || 0);
    return telenor + jazz + ufone + zong;
});

const stampPaperTotal = computed(() => {
    return sheet.stampPaper.reduce((sum, item) => sum + (Number(item.amount) || 0), 0);
});

const recoveryTotal = computed(() => {
    const recoverySum = sheet.recovery.reduce(
        (sum, item) => sum + (Number(item.amount) || 0),
        0
    );
    const stampPaperSum = sheet.stampPaper.reduce(
        (sum, item) => sum + (Number(item.amount) || 0),
        0
    );
    return recoverySum + stampPaperSum;
});

const getTotalSending = (
    entries: { sending: string; receiving: string }[]
) => {
    return entries.reduce((sum, entry) => {
        return sum + (isPureNumber(entry.sending)
            ? Number(entry.sending)
            : 0);
    }, 0);
};

const getTotalReceiving = (
    entries: { sending: string; receiving: string }[]
) => {
    return entries.reduce((sum, entry) => {
        return sum + (isPureNumber(entry.receiving)
            ? Number(entry.receiving)
            : 0);
    }, 0);
};

const extractLastBalance = (
    entries: { sending: string; receiving: string }[]
) => {
    const regex = /b\s+(\d+)/i;

    for (let i = entries.length - 1; i >= 0; i--) {
        const sendMatch = regex.exec(entries[i]?.sending || "");

        if (sendMatch) {
            return Number(sendMatch[1]);
        }

        const receiveMatch = regex.exec(entries[i]?.receiving || "");

        if (receiveMatch) {
            return Number(receiveMatch[1]);
        }
    }

    return null;
};

const purchasingTotal = computed(() => {
    return (
        getTotalReceiving(sheet.omni) +
        getTotalReceiving(sheet.easypaisa) +
        getTotalReceiving(sheet.jazzcash) +
        getTotalReceiving(sheet.epaccount) +
        getTotalReceiving(sheet.jcaccount) +
        sheet.manualpurchasing.reduce(
            (sum, item) => sum + (Number(item.amount) || 0),
            0
        ) +
        redBookTotal.value
    );
});

const cashInfoTotal = computed(() => {
    return (
        getTotalSending(sheet.omni) +
        getTotalSending(sheet.easypaisa) +
        getTotalSending(sheet.jazzcash) +
        getTotalSending(sheet.epaccount) +
        getTotalSending(sheet.jcaccount) +
        recoveryTotal.value +
        totalELoad.value +
        n(sheet.extra)
    );
});

const redBookTotal = computed(() => {
    return sheet.redBook.reduce(
        (sum, item) => sum + (Number(item.amount) || 0),
        0
    );
});

const cashTotal = computed(() => {
    return (
        5000 * n(sheet.cash5000) +
        1000 * n(sheet.cash1000) +
        500 * n(sheet.cash500) +
        100 * n(sheet.cash100) +
        50 * n(sheet.cash50) +
        20 * n(sheet.cash20) +
        10 * n(sheet.cash10) +
        5 * n(sheet.cash5)
    );
});

const handleDownloadPDF = async () => {
    try {
        const s = sheet;
        const n = (v: any) => Number(v) || 0;

        const today = new Date();
        const day = String(today.getDate()).padStart(2, "0");
        const month = String(today.getMonth() + 1).padStart(2, "0");
        const year = today.getFullYear();

        const fileName = `KMKCommunicationSheet-${day}-${month}-${year}.pdf`;

        const docDefinition: any = {
            pageSize: "A4",
            pageOrientation: "portrait",
            pageMargins: [20, 25, 20, 25],

            defaultStyle: {
                fontSize: 7.2
            },

            styles: {
                header: { fontSize: 12, bold: true, alignment: "center", margin: [0, 0, 0, 2], decoration: "underline" },
                subheader: { fontSize: 8, bold: true, alignment: "center", decoration: "underline", margin: [0, 4, 0, 3] },
                tableHeader: { bold: true, fillColor: "#e0e0e0", fontSize: 7 },
                sectionHeader: {
                    bold: true,
                    fontSize: 7,
                    fillColor: "#d3d3d3",
                    color: "#000000",
                    alignment: "center"
                },
                small: { fontSize: 7 },
            },

            content: [
                { text: "KMK COMMUNICATION", style: "header" },

                {
                    columns: [
                        { text: `Date: ${dateTime.value.date}`, alignment: "left", style: "small" },
                        {
                            text: dateTime.value.pdfIslamicDate,
                            alignment: "center",
                            fontSize: 10,
                            font: "Urdu",
                        },
                        { text: `Time: ${dateTime.value.time}`, alignment: "right", style: "small" },
                    ],
                    margin: [0, 0, 0, 3],
                },

                { text: "EASYLOAD & ACCOUNTS", style: "subheader" },

                {
                    columns: [
                        {
                            width: "48%",
                            table: {
                                widths: ["22%", "19%", "19%", "19%", "21%"],
                                body: [
                                    [{ text: "EasyLoad", colSpan: 5, style: "tableHeader", alignment: "center" }, {}, {}, {}, {}],
                                    ["", "Telenor", "Jazz", "Ufone", "Zong"],
                                    ["Opening", n(s.telenorOpeningBalance), n(s.jazzOpeningBalance), n(s.ufoneOpeningBalance), n(s.zongOpeningBalance)],
                                    ["New", n(s.telenorNewBalance), n(s.jazzNewBalance), n(s.ufoneNewBalance), n(s.zongNewBalance)],
                                    ["Rev", n(s.telenorReversalBalance), n(s.jazzReversalBalance), n(s.ufoneReversalBalance), n(s.zongReversalBalance)],
                                    ["Total",
                                        n(s.telenorOpeningBalance) + n(s.telenorNewBalance) + n(s.telenorReversalBalance),
                                        n(s.jazzOpeningBalance) + n(s.jazzNewBalance) + n(s.jazzReversalBalance),
                                        n(s.ufoneOpeningBalance) + n(s.ufoneNewBalance) + n(s.ufoneReversalBalance),
                                        n(s.zongOpeningBalance) + n(s.zongNewBalance) + n(s.zongReversalBalance)
                                    ],
                                    ["Closing", n(s.telenorClosingBalance), n(s.jazzClosingBalance), n(s.ufoneClosingBalance), n(s.zongClosingBalance)],
                                    ["Sell", getELoadSell("telenor"), getELoadSell("jazz"), getELoadSell("ufone"), getELoadSell("zong")],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },

                        { width: "2%", text: "" },

                        {
                            width: "50%",
                            stack: [
                                {
                                    table: {
                                        widths: ["30%", "35%", "35%"],
                                        body: [
                                            [{ text: "Bank Accounts", colSpan: 3, style: "sectionHeader", alignment: "center" }, {}, {}],
                                            ["Account No.", "265999891", "37300247"],
                                            ["Balance", n(s.accountBalance265999891), n(s.accountBalance37300247)],
                                            ["Deposit", n(s.deposit265999891), n(s.deposit37300247)],
                                            ["Withdraw", n(s.withdrawl265999891), n(s.withdrawl37300247)],
                                            ["Remaining Balance",
                                                n(s.accountBalance265999891) + n(s.deposit265999891) - n(s.withdrawl265999891),
                                                n(s.accountBalance37300247) + n(s.deposit37300247) - n(s.withdrawl37300247)
                                            ],
                                        ],
                                    },
                                    layout: "lightHorizontalLines",
                                },
                            ],
                        },
                    ],
                },

                { text: "\nRECOVERY & PURCHASING", style: "subheader" },

                {
                    columns: [
                        {
                            width: "20%",
                            table: {
                                widths: ["65%", "35%"],
                                body: [
                                    [{ text: "Stamp Paper", colSpan: 2, style: "tableHeader", alignment: "center" }, {}],
                                    ...sheet.stampPaper
                                        .filter(r => r.amount != null && r.amount > 0)
                                        .map(r => ["Stamp Paper", n(r.amount)]),
                                    ["Total", { text: stampPaperTotal.value, bold: true }],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },

                        { width: "1%", text: "" },

                        {
                            width: "25%",
                            table: {
                                widths: ["65%", "35%"],
                                body: [
                                    [{ text: "Recovery / Sell", colSpan: 2, style: "tableHeader", alignment: "center" }, {}],
                                    ...sheet.recovery
                                        .filter(r => r.name || r.amount != null)
                                        .map(r => [r.name || "-", n(r.amount)]),
                                    ["Total", { text: recoveryTotal.value, bold: true }],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },

                        { width: "1%", text: "" },

                        {
                            width: "28%",
                            table: {
                                widths: ["65%", "35%"],
                                body: [
                                    [{ text: "Home Purchasing", colSpan: 2, style: "tableHeader", alignment: "center" }, {}],
                                    ...sheet.redBook.filter(r => r.name || r.amount != null).map(r => [r.name || "-", n(r.amount)]),
                                    ["Total", { text: redBookTotal.value, bold: true }],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },

                        { width: "1%", text: "" },

                        {
                            width: "25%",
                            table: {
                                widths: ["65%", "35%"],
                                body: [
                                    [{ text: "Purchasing", colSpan: 2, style: "tableHeader", alignment: "center" }, {}],
                                    ["Omni/EP/JC Rec", getTotalReceiving(sheet.omni) + getTotalReceiving(sheet.easypaisa) + getTotalReceiving(sheet.jazzcash)],
                                    ["EP/JC Acc", getTotalReceiving(sheet.epaccount) + getTotalReceiving(sheet.jcaccount)],
                                    ...sheet.manualpurchasing
                                        .filter((item) => item.name || item.amount != null)
                                        .map((item, index) => [item.name || `Manual ${index + 1}`, n(item.amount)]),
                                    ["Home Purchasing", redBookTotal.value],
                                    ["Total", { text: purchasingTotal.value, bold: true }],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },
                    ],
                },

                { text: "\nSTATEMENTS", style: "subheader" },

                {
                    columns: [
                        { width: "20%", stack: [{ text: `UBL Omni ${sheet.lastBalances.omni ? `(Op Bal - ${sheet.lastBalances.omni})` : ""}`, style: "sectionHeader", margin: [0, 0, 0, 4] }, { table: { widths: ["*", "*"], body: [["Sending", "Receiving"], ...sheet.omni.filter(r => r.sending || r.receiving).map(r => [r.sending || "", r.receiving || ""]), ["Total Send", getTotalSending(sheet.omni)], ["Total Rec", getTotalReceiving(sheet.omni)], ["Last Bal", extractLastBalance(sheet.omni) || "-"]] }, layout: "lightHorizontalLines" }] },
                        { width: "20%", stack: [{ text: `EasyPaisa ${sheet.lastBalances.easypaisa ? `(Op Bal - ${sheet.lastBalances.easypaisa})` : ""}`, style: "sectionHeader", margin: [0, 0, 0, 4] }, { table: { widths: ["*", "*"], body: [["Sending", "Receiving"], ...sheet.easypaisa.filter(r => r.sending || r.receiving).map(r => [r.sending || "", r.receiving || ""]), ["Total Send", getTotalSending(sheet.easypaisa)], ["Total Rec", getTotalReceiving(sheet.easypaisa)], ["Last Bal", extractLastBalance(sheet.easypaisa) || "-"]] }, layout: "lightHorizontalLines" }] },
                        { width: "20%", stack: [{ text: `JazzCash ${sheet.lastBalances.jazzcash ? `(Op Bal - ${sheet.lastBalances.jazzcash})` : ""}`, style: "sectionHeader", margin: [0, 0, 0, 4] }, { table: { widths: ["*", "*"], body: [["Sending", "Receiving"], ...sheet.jazzcash.filter(r => r.sending || r.receiving).map(r => [r.sending || "", r.receiving || ""]), ["Total Send", getTotalSending(sheet.jazzcash)], ["Total Rec", getTotalReceiving(sheet.jazzcash)], ["Last Bal", extractLastBalance(sheet.jazzcash) || "-"]] }, layout: "lightHorizontalLines" }] },
                        { width: "20%", stack: [{ text: `EP Account ${sheet.lastBalances.epaccount ? `(Op Bal - ${sheet.lastBalances.epaccount})` : ""}`, style: "sectionHeader", margin: [0, 0, 0, 4] }, { table: { widths: ["*", "*"], body: [["Sending", "Receiving"], ...sheet.epaccount.filter(r => r.sending || r.receiving).map(r => [r.sending || "", r.receiving || ""]), ["Total Send", getTotalSending(sheet.epaccount)], ["Total Rec", getTotalReceiving(sheet.epaccount)], ["Last Bal", extractLastBalance(sheet.epaccount) || "-"]] }, layout: "lightHorizontalLines" }] },
                        { width: "20%", stack: [{ text: `JC Merchant ${sheet.lastBalances.jcaccount ? `(Op Bal - ${sheet.lastBalances.jcaccount})` : ""}`, style: "sectionHeader", margin: [0, 0, 0, 4] }, { table: { widths: ["*", "*"], body: [["Sending", "Receiving"], ...sheet.jcaccount.filter(r => r.sending || r.receiving).map(r => [r.sending || "", r.receiving || ""]), ["Total Send", getTotalSending(sheet.jcaccount)], ["Total Rec", getTotalReceiving(sheet.jcaccount)], ["Last Bal", extractLastBalance(sheet.jcaccount) || "-"]] }, layout: "lightHorizontalLines" }] },
                    ],
                    columnGap: 10,
                },

                { text: "\nCASH SUMMARY", style: "subheader" },

                {
                    columns: [
                        {
                            width: "33%",
                            table: {
                                widths: ["28%", "22%", "8%", "42%"],
                                body: [
                                    [{ text: "Cash Detail", colSpan: 4, style: "tableHeader", alignment: "center" }, {}, {}, {}],
                                    ["Denom", "Qty", "=", "Amount"],
                                    ["5000", n(s.cash5000), "=", 5000 * n(s.cash5000)],
                                    ["1000", n(s.cash1000), "=", 1000 * n(s.cash1000)],
                                    ["500", n(s.cash500), "=", 500 * n(s.cash500)],
                                    ["100", n(s.cash100), "=", 100 * n(s.cash100)],
                                    ["50", n(s.cash50), "=", 50 * n(s.cash50)],
                                    ["20", n(s.cash20), "=", 20 * n(s.cash20)],
                                    ["10", n(s.cash10), "=", 10 * n(s.cash10)],
                                    ["5", n(s.cash5), "=", 5 * n(s.cash5)],
                                    [{ text: "TOTAL", bold: true, colSpan: 3 }, {}, {}, { text: cashTotal.value, bold: true }],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },

                        { width: "1%", text: "" },

                        {
                            width: "33%",
                            table: {
                                widths: ["50%", "50%"],
                                body: [
                                    [{ text: "Analytics", colSpan: 2, style: "tableHeader", alignment: "center" }, {}],
                                    ["Previous Cash", n(s.previousCash)],
                                    ["Today Cash", cashInfoTotal.value],
                                    ["Total Cash", n(s.previousCash) + cashInfoTotal.value],
                                    ["Purchasing", purchasingTotal.value],
                                    ["Remaining", (n(s.previousCash) + cashInfoTotal.value) - purchasingTotal.value],
                                    [{ text: "DIFFERENCE", bold: true, fontSize: 8 }, { text: (cashTotal.value - ((n(s.previousCash) + cashInfoTotal.value) - purchasingTotal.value)), bold: true, fontSize: 8 }],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },

                        { width: "1%", text: "" },

                        {
                            width: "33%",
                            table: {
                                widths: ["65%", "35%"],
                                body: [
                                    [{ text: "Cash Info", colSpan: 2, style: "tableHeader", alignment: "center" }, {}],
                                    ["UBL Omni Sending", getTotalSending(sheet.omni)],
                                    ["EasyPaisa Sending", getTotalSending(sheet.easypaisa)],
                                    ["JazzCash Sending", getTotalSending(sheet.jazzcash)],
                                    ["EP Account Sending", getTotalSending(sheet.epaccount)],
                                    ["JC Account Sending", getTotalSending(sheet.jcaccount)],
                                    ["Recovery", recoveryTotal.value],
                                    ["EasyLoad", totalELoad.value],
                                    ["Extra", n(s.extra)],
                                    ["Total", cashInfoTotal.value],
                                ],
                            },
                            layout: "lightHorizontalLines",
                        },
                    ],
                    columnGap: 3,
                },
            ],
        };

        const pdf = pdfMake.createPdf(docDefinition);
        const base64 = await pdf.getBase64();

        const byteCharacters = atob(base64);
        const byteNumbers = new Array(byteCharacters.length);
        for (let i = 0; i < byteCharacters.length; i++) {
            byteNumbers[i] = byteCharacters.charCodeAt(i);
        }
        const pdfArray = Array.from(new Uint8Array(byteNumbers));

        const jsonData = JSON.stringify(sheet, null, 2);

        await SavePDF(fileName, pdfArray, jsonData);

        toast.success(`File: ${fileName} PDF generated and saved successfully!`);
        toast.success(`JSON Backup saved successfully!`);
    } catch (err: any) {
        console.error("PDF Error:", err);
        toast.error(`Error generating PDF: ${err.message || err}`);
    }
};

const clearLocalStorage = () => {
    localStorage.clear();
    toast.success("Local storage cleared");
};
</script>

<template>
    <div id="print-area">
        <div class="container">
            <div v-if="DEBUG">
                <pre v-for="([key, value], index) in Object.entries(sheet)" :key="index">
                    <strong>{{ key }}:</strong>
                    {{ typeof value === 'object' ? JSON.stringify(value, null, 2) : value }}
                </pre>
            </div>

            <div class="header">
                <div class="title-container">
                    <h1>KMK Communication</h1>
                </div>
                <div class="time-actions-container">
                    <div>
                        <h3 id="current-date">{{ dateTime.date }}</h3>
                        <h3 id="current-islamic-date" dir="rtl" style="font-family: Urdu;">
                            {{ dateTime.islamicDate }}
                        </h3>
                    </div>
                    <h3 id="current-time">{{ dateTime.time }}</h3>
                    <div class="actions no-print">
                        <button class="button" @click="showCashModal = true">
                            Cash Calculator
                        </button>
                        <button class="button" @click="handleDownloadPDF">
                            Save PDF
                        </button>
                        <button class="button" @click="uploadBackup">
                            Upload Backup
                        </button>
                        <button class="button-danger" @click="showClearSheetModal = true">
                            Clear Sheet
                        </button>
                        <button v-if="DEBUG" class="button-danger" @click="clearLocalStorage">
                            Clear Local Storage
                        </button>
                    </div>
                </div>
            </div>

            <div class="eload-accounts-container">
                <div>
                    <div class="easyload-container">
                        <div class="easyload">
                            <table>
                                <tbody>
                                    <tr>
                                        <th />
                                        <th>Telenor</th>
                                        <th>Jazz</th>
                                        <th>Ufone</th>
                                        <th>Zong</th>
                                    </tr>

                                    <tr>
                                        <td>Opening Balance</td>
                                        <td><input v-model.number="sheet.telenorOpeningBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.jazzOpeningBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.ufoneOpeningBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.zongOpeningBalance" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>New Balance</td>
                                        <td><input v-model.number="sheet.telenorNewBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.jazzNewBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.ufoneNewBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.zongNewBalance" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Reversal Balance</td>
                                        <td><input v-model.number="sheet.telenorReversalBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.jazzReversalBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.ufoneReversalBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.zongReversalBalance" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Total Rs</td>
                                        <td><input
                                                :value="n(sheet.telenorOpeningBalance) + n(sheet.telenorNewBalance) + n(sheet.telenorReversalBalance)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="n(sheet.jazzOpeningBalance) + n(sheet.jazzNewBalance) + n(sheet.jazzReversalBalance)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="n(sheet.ufoneOpeningBalance) + n(sheet.ufoneNewBalance) + n(sheet.ufoneReversalBalance)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="n(sheet.zongOpeningBalance) + n(sheet.zongNewBalance) + n(sheet.zongReversalBalance)"
                                                type="number" disabled /></td>
                                    </tr>

                                    <tr>
                                        <td>Closing Balance</td>
                                        <td><input v-model.number="sheet.telenorClosingBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.jazzClosingBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.ufoneClosingBalance" type="number" /></td>
                                        <td><input v-model.number="sheet.zongClosingBalance" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Total ELoad Sell</td>
                                        <td><input :value="getELoadSell('telenor')" type="number" disabled /></td>
                                        <td><input :value="getELoadSell('jazz')" type="number" disabled /></td>
                                        <td><input :value="getELoadSell('ufone')" type="number" disabled /></td>
                                        <td><input :value="getELoadSell('zong')" type="number" disabled /></td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>

                    <div class="accounts-container space-y">
                        <div class="accounts">
                            <table>
                                <tbody>
                                    <tr>
                                        <th>Accounts</th>
                                        <th>265999891</th>
                                        <th>37300247</th>
                                    </tr>

                                    <tr>
                                        <td>Rs</td>
                                        <td><input v-model="sheet.accountBalance265999891" type="number" /></td>
                                        <td><input v-model="sheet.accountBalance37300247" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Deposit</td>
                                        <td><input v-model="sheet.deposit265999891" type="number" /></td>
                                        <td><input v-model="sheet.deposit37300247" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Total Rs</td>
                                        <td><input :value="n(sheet.accountBalance265999891) + n(sheet.deposit265999891)"
                                                type="number" disabled /></td>
                                        <td><input :value="n(sheet.accountBalance37300247) + n(sheet.deposit37300247)"
                                                type="number" disabled /></td>
                                    </tr>

                                    <tr>
                                        <td>Withdrawl</td>
                                        <td><input v-model="sheet.withdrawl265999891" type="number" /></td>
                                        <td><input v-model="sheet.withdrawl37300247" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Remaining Balance</td>
                                        <td><input
                                                :value="(n(sheet.accountBalance265999891) + n(sheet.deposit265999891)) - n(sheet.withdrawl265999891)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="(n(sheet.accountBalance37300247) + n(sheet.deposit37300247)) - n(sheet.withdrawl37300247)"
                                                type="number" disabled /></td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                        <div class="eload-summery mx">
                            <table>
                                <tbody>
                                    <tr>
                                        <th colspan="2">Easyload Summery</th>
                                    </tr>
                                    <tr>
                                        <td>Telenor</td>
                                        <td><input :value="getELoadSell('telenor')" type="number" disabled /></td>
                                    </tr>
                                    <tr>
                                        <td>Jazz</td>
                                        <td><input :value="getELoadSell('jazz')" type="number" disabled /></td>
                                    </tr>
                                    <tr>
                                        <td>Ufone</td>
                                        <td><input :value="getELoadSell('ufone')" type="number" disabled /></td>
                                    </tr>
                                    <tr>
                                        <td>Zong</td>
                                        <td><input :value="getELoadSell('zong')" type="number" disabled /></td>
                                    </tr>
                                    <tr>
                                        <td>Total</td>
                                        <td><input :value="totalELoad" type="number" disabled /></td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>

                <div class="purchasing-recovery-container">
                    <div class="stamp-paper-container">
                        <table>
                            <tbody>
                                <tr>
                                    <th colspan="1">Stamp Paper</th>
                                </tr>
                                <tr v-for="(row, index) in sheet.stampPaper" :key="index">
                                    <td>
                                        <input type="number" v-model.number="row.amount" />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <div class="recovery-container">
                        <table>
                            <tbody>
                                <tr>
                                    <th colspan="2">Recovery/Sell</th>
                                </tr>

                                <tr v-for="(row, index) in sheet.recovery" :key="index">
                                    <td>
                                        <input type="text" v-model="row.name" />
                                    </td>
                                    <td>
                                        <input type="number" v-model.number="row.amount" />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Total</td>
                                    <td>
                                        <input type="number" :value="recoveryTotal" disabled />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                    <div class="purchasing-container">
                        <table>
                            <tbody>
                                <tr>
                                    <th colspan="2">Purchasing/Borrow</th>
                                </tr>
                                <tr>
                                    <td>Omni/EP/JC Rec</td>
                                    <td>
                                        <input
                                            :value="getTotalReceiving(sheet.omni) + getTotalReceiving(sheet.easypaisa) + getTotalReceiving(sheet.jazzcash)"
                                            type="number" disabled />
                                    </td>
                                </tr>
                                <tr>
                                    <td>EP/JC Account</td>
                                    <td>
                                        <input
                                            :value="getTotalReceiving(sheet.epaccount) + getTotalReceiving(sheet.jcaccount)"
                                            type="number" disabled />
                                    </td>
                                </tr>
                                <tr v-for="(row, index) in sheet.manualpurchasing" :key="index">
                                    <td>
                                        <input type="text" v-model="row.name" />
                                    </td>
                                    <td>
                                        <input type="number" v-model.number="row.amount" />
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <input type="text" value="Home Prchas" disabled />
                                    </td>
                                    <td>
                                        <input :value="redBookTotal" type="number" disabled />
                                    </td>
                                </tr>
                                <tr>
                                    <td>Total Purchsing</td>
                                    <td>
                                        <input :value="purchasingTotal" type="number" disabled />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>

            <div class="statements-container">
                <div class="ubl-omni">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">UBL Omni</th>
                            </tr>
                            <tr>
                                <td>Balance</td>
                                <td>
                                    <input v-model.number="sheet.lastBalances.omni" type="number" />
                                </td>
                            </tr>
                            <tr>
                                <td>Sending</td>
                                <td>Receiving</td>
                            </tr>
                            <tr v-for="(row, index) in sheet.omni" :key="index">
                                <td>
                                    <input type="text" v-model="row.sending" />
                                </td>
                                <td>
                                    <input type="text" v-model="row.receiving" />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.omni)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Receiving</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.omni)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Last Balance</td>
                                <td>
                                    <input :value="extractLastBalance(sheet.omni)" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="easypaisa">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">Easy Paisa</th>
                            </tr>

                            <tr>
                                <td>Balance</td>
                                <td>
                                    <input v-model.number="sheet.lastBalances.easypaisa" type="number" />
                                </td>
                            </tr>

                            <tr>
                                <td>Sending</td>
                                <td>Receiving</td>
                            </tr>

                            <tr v-for="(row, index) in sheet.easypaisa" :key="index">
                                <td>
                                    <input type="text" v-model="row.sending" />
                                </td>

                                <td>
                                    <input type="text" v-model="row.receiving" />
                                </td>
                            </tr>

                            <tr>
                                <td>Total Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.easypaisa)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Receiving</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.easypaisa)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Last Balance</td>
                                <td>
                                    <input :value="extractLastBalance(sheet.easypaisa)" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="jazzcash">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">JazzCash</th>
                            </tr>

                            <tr>
                                <td>Balance</td>
                                <td>
                                    <input v-model.number="sheet.lastBalances.jazzcash" type="number" />
                                </td>
                            </tr>

                            <tr>
                                <td>Sending</td>
                                <td>Receiving</td>
                            </tr>

                            <tr v-for="(row, index) in sheet.jazzcash" :key="index">
                                <td>
                                    <input type="text" v-model="row.sending" />
                                </td>

                                <td>
                                    <input type="text" v-model="row.receiving" />
                                </td>
                            </tr>

                            <tr>
                                <td>Total Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.jazzcash)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Receiving</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.jazzcash)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Last Balance</td>
                                <td>
                                    <input :value="extractLastBalance(sheet.jazzcash)" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="ep-account">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">EP Account 0333</th>
                            </tr>

                            <tr>
                                <td>Balance</td>
                                <td>
                                    <input v-model.number="sheet.lastBalances.epaccount" type="number" />
                                </td>
                            </tr>

                            <tr>
                                <td>Sending</td>
                                <td>Receiving</td>
                            </tr>

                            <tr v-for="(row, index) in sheet.epaccount" :key="index">
                                <td>
                                    <input type="text" v-model="row.sending" />
                                </td>

                                <td>
                                    <input type="text" v-model="row.receiving" />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.epaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Receiving</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.epaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Last Balance</td>
                                <td>
                                    <input :value="extractLastBalance(sheet.epaccount)" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="jc-account">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">JC Merchant Account</th>
                            </tr>

                            <tr>
                                <td>Balance</td>
                                <td>
                                    <input v-model.number="sheet.lastBalances.jcaccount" type="number" />
                                </td>
                            </tr>

                            <tr>
                                <td>Sending</td>
                                <td>Receiving</td>
                            </tr>

                            <tr v-for="(row, index) in sheet.jcaccount" :key="index">
                                <td>
                                    <input type="text" v-model="row.sending" />
                                </td>
                                <td>
                                    <input type="text" v-model="row.receiving" />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.jcaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Receiving</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.jcaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Last Balance</td>
                                <td>
                                    <input :value="extractLastBalance(sheet.jcaccount)" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div class="analytics-container">
                <div class="cash-detail">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="4">Cash Detail</th>
                            </tr>
                            <tr>
                                <td>5000</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash5000" type="number" /></td>
                                <td><input :value="5000 * n(sheet.cash5000)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>1000</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash1000" type="number" /></td>
                                <td><input :value="1000 * n(sheet.cash1000)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>500</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash500" type="number" /></td>
                                <td><input :value="500 * n(sheet.cash500)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>100</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash100" type="number" /></td>
                                <td><input :value="100 * n(sheet.cash100)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>50</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash50" type="number" /></td>
                                <td><input :value="50 * n(sheet.cash50)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>20</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash20" type="number" /></td>
                                <td><input :value="20 * n(sheet.cash20)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>10</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash10" type="number" /></td>
                                <td><input :value="10 * n(sheet.cash10)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td>5</td>
                                <td>x</td>
                                <td><input v-model="sheet.cash5" type="number" /></td>
                                <td><input :value="5 * n(sheet.cash5)" type="number" disabled /></td>
                            </tr>

                            <tr>
                                <td colspan="3">Total</td>
                                <td>
                                    <input :value="cashTotal" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="analytics">
                    <table>
                        <tbody>
                            <tr>
                                <td>Previous Cash</td>
                                <td>
                                    <input v-model.number="sheet.previousCash" type="number" />
                                </td>
                            </tr>
                            <tr>
                                <td>Today Cash</td>
                                <td>
                                    <input :value="cashInfoTotal" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Total Today & Previous Cash</td>
                                <td>
                                    <input :value="n(sheet.previousCash) + cashInfoTotal" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Purchasing</td>
                                <td>
                                    <input :value="purchasingTotal" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Remaining Cash</td>
                                <td>
                                    <input :value="(n(sheet.previousCash) + cashInfoTotal) - purchasingTotal"
                                        type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td style="font-size: 1.5rem;">Difference</td>
                                <td>
                                    <input style="font-size: 1.1rem;"
                                        :value="(cashTotal - ((n(sheet.previousCash) + cashInfoTotal) - purchasingTotal))"
                                        type="number" readonly />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="red-book">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">HOME Purchasing</th>
                            </tr>
                            <tr v-for="(row, index) in sheet.redBook" :key="index">
                                <td>
                                    <input type="text" v-model="row.name" />
                                </td>
                                <td>
                                    <input type="number" v-model.number="row.amount" />
                                </td>
                            </tr>
                            <tr>
                                <td>Total</td>
                                <td>
                                    <input type="number" :value="redBookTotal" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                <div class="cash-info">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">Cash Info</th>
                            </tr>
                            <tr>
                                <td>UBL Omni Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.omni)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>EasyPaisa Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.easypaisa)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>JazzCash Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.jazzcash)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>EP Account Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.epaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>JC M Account Sending</td>
                                <td>
                                    <input :value="getTotalSending(sheet.jcaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Recovery</td>
                                <td>
                                    <input :value="recoveryTotal" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>EasyLoad</td>
                                <td>
                                    <input :value="totalELoad" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Extra</td>
                                <td>
                                    <input v-model.number="sheet.extra" type="number" />
                                </td>
                            </tr>
                            <tr>
                                <td>Total</td>
                                <td>
                                    <input :value="cashInfoTotal" type="number" disabled />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>

    <div v-if="showClearSheetModal" class="modal-overlay">
        <div class="modal-box">
            <h2>⚠️ Clear Sheet ⚠️</h2>

            <p>
                Are you sure you want to clear the sheet?
            </p>

            <div class="modal-actions">
                <button class="button" @click="showClearSheetModal = false">
                    Cancel
                </button>

                <button class="button-danger" @click="confirmClearSheet">
                    Yes, Clear It!
                </button>
            </div>
        </div>
    </div>
    <CashModal v-model="showCashModal" />
    <Toaster position="top-center" richColors expand :visible-toasts="10" />
</template>

<style scoped>
.modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
}

.modal-box {
    background: white;
    padding: 12px;
    border-radius: 12px;
    width: 500px;
    max-width: 90%;
    box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
}

.modal-box h2 {
    text-align: center;
}

.modal-box p {
    font-weight: 500;
    text-align: center;
    font-size: 1.2rem;
}

.modal-actions {
    display: flex;
    align-items: center;
    justify-content: space-evenly;
}

.modal-actions button {
    font-size: 1rem;
}
</style>
