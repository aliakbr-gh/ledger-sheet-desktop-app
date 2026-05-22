<script setup lang="ts">
import { ref, onMounted, onUnmounted, reactive, watch, computed } from "vue";
import jsPDF from "jspdf";
import html2canvas from "html2canvas";
import { SavePDF } from "../wailsjs/go/main/App";

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
    accountBalance266001445: number | null;
    accountBalance37300247: number | null;
    accountBalance257283991: number | null;

    borrow: { name: string; amount: number | null }[];
    recovery: { name: string; amount: number | null }[];

    deposit265999891: number | null;
    deposit266001445: number | null;
    deposit37300247: number | null;
    deposit257283991: number | null;

    withdrawl265999891: number | null;
    withdrawl266001445: number | null;
    withdrawl37300247: number | null;
    withdrawl257283991: number | null;

    totalCards: number | null;
    sellCards: number | null;

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
    accountBalance266001445: null,
    accountBalance37300247: null,
    accountBalance257283991: null,

    borrow: Array.from({ length: 12 }, () => ({
        name: "",
        amount: null as number | null,
    })),

    recovery: Array.from({ length: 12 }, () => ({
        name: "",
        amount: null as number | null,
    })),

    deposit265999891: null,
    deposit266001445: null,
    deposit37300247: null,
    deposit257283991: null,

    withdrawl265999891: null,
    withdrawl266001445: null,
    withdrawl37300247: null,
    withdrawl257283991: null,

    totalCards: null,
    sellCards: null,

    omni: Array.from({ length: 11 }, () => ({
        sending: "",
        receiving: "",
    })),

    easypaisa: Array.from({ length: 11 }, () => ({
        sending: "",
        receiving: "",
    })),

    jazzcash: Array.from({ length: 11 }, () => ({
        sending: "",
        receiving: "",
    })),

    epaccount: Array.from({ length: 11 }, () => ({
        sending: "",
        receiving: "",
    })),

    jcaccount: Array.from({ length: 11 }, () => ({
        sending: "",
        receiving: "",
    })),

    manualpurchasing: Array.from({ length: 10 }, () => ({
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

const savedSheet = localStorage.getItem("sheet");

if (savedSheet) {
    const parsed = JSON.parse(savedSheet);

    Object.assign(sheet, parsed);

    if (!parsed.borrow) {
        sheet.borrow = Array.from({ length: 12 }, () => ({
            name: "",
            amount: null,
        }));
    }

    if (!parsed.recovery) {
        sheet.recovery = Array.from({ length: 12 }, () => ({
            name: "",
            amount: null,
        }));
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

const clearSheet = () => {

    sheet.previousCash = cashTotal.value;

    sheet.telenorOpeningBalance = sheet.telenorClosingBalance;
    sheet.jazzOpeningBalance = sheet.jazzClosingBalance;
    sheet.ufoneOpeningBalance = sheet.ufoneClosingBalance;
    sheet.zongOpeningBalance = sheet.zongClosingBalance;

    sheet.telenorNewBalance = null;
    sheet.jazzNewBalance = null;
    sheet.ufoneNewBalance = null;
    sheet.zongNewBalance = null;

    sheet.telenorReversalBalance = null,
    sheet.jazzReversalBalance = null,
    sheet.ufoneReversalBalance = null,
    sheet.zongReversalBalance = null,

    sheet.accountBalance265999891 = null,
    sheet.accountBalance266001445 = null,
    sheet.accountBalance37300247 = null,
    sheet.accountBalance257283991 = null,

    sheet.deposit265999891 = null,
    sheet.deposit266001445 = null,
    sheet.deposit37300247 = null,
    sheet.deposit257283991 = null,

    sheet.withdrawl265999891 = null,
    sheet.withdrawl266001445 = null,
    sheet.withdrawl37300247 = null,
    sheet.withdrawl257283991 = null,

    sheet.totalCards = null,
    sheet.sellCards = null,

    sheet.lastBalances.omni = extractLastBalance(sheet.omni);
    sheet.lastBalances.easypaisa = extractLastBalance(sheet.easypaisa);
    sheet.lastBalances.jazzcash = extractLastBalance(sheet.jazzcash);
    sheet.lastBalances.epaccount = extractLastBalance(sheet.epaccount);
    sheet.lastBalances.jcaccount = extractLastBalance(sheet.jcaccount);

    const shift = (type: "telenor" | "jazz" | "ufone" | "zong") => {
        const closing = sheet[`${type}ClosingBalance`];

        if (closing !== null && closing !== undefined) {
            sheet[`${type}OpeningBalance`] = closing;
        }
    };

    shift("telenor");
    shift("jazz");
    shift("ufone");
    shift("zong");

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

    sheet.borrow.forEach((item) => {
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
};

// Helper Functions
const n = (v: number | null | undefined) => v ?? 0;

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
    time: "",
});

const updateDateTime = () => {
    const now = new Date();

    dateTime.value.date = now.toLocaleDateString("en-US", {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric",
    });

    dateTime.value.time = now.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
    });
};

let interval: number;

onMounted(() => {
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

const borrowTotal = computed(() => {
    return sheet.borrow.reduce(
        (sum, item) => sum + (Number(item.amount) || 0),
        0
    );
});

const recoveryTotal = computed(() => {
    return sheet.recovery.reduce(
        (sum, item) => sum + (Number(item.amount) || 0),
        0
    );
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
        )
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
        (100 * n(sheet.sellCards)) +
        totalELoad.value +
        n(sheet.extra)
    );
});

const manualPurchasingTotal = computed(() => {
    return sheet.manualpurchasing.reduce(
        (sum, item) => sum + (Number(item.amount) || 0),
        0
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
    const element = document.getElementById("print-area");
    if (!element) return;

    const noPrintEls = document.querySelectorAll(".no-print");

    noPrintEls.forEach((el) => {
        (el as HTMLElement).style.display = "none";
    });

    await new Promise(resolve => setTimeout(resolve, 100));

    try {
        const canvas = await html2canvas(element, {
            scale: 2,
            useCORS: true,
            backgroundColor: "#ffffff",
        });

        noPrintEls.forEach((el) => {
            (el as HTMLElement).style.display = "";
        });

        const imgData = canvas.toDataURL("image/jpeg", 1.0);

        const pdf = new jsPDF({
            orientation: "landscape",
            unit: "mm",
            format: "a4",
        });

        const pageWidth = 297;
        const imgHeight = (canvas.height * pageWidth) / canvas.width;


        const now = new Date();

        const day = String(now.getDate()).padStart(2, "0");

        const month = String(now.getMonth() + 1).padStart(2, "0");

        const year = now.getFullYear();

        pdf.addImage(imgData, "JPEG", 0, 0, pageWidth, imgHeight);

        const fileName = `KMKCommunicationSheet-${day}-${month}-${year}.pdf`;

        const pdfArrayBuffer = pdf.output("arraybuffer");

        await SavePDF(
            fileName,
            Array.from(new Uint8Array(pdfArrayBuffer))
        );

    } catch (error) {
        console.error(error);

        noPrintEls.forEach((el) => {
            (el as HTMLElement).style.display = "";
        });
    }
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
                    <h2 id="current-date">{{ dateTime.date }}</h2>
                    <h2 id="current-time">{{ dateTime.time }}</h2>
                    <div class="actions no-print">
                        <button class="button" @click="handleDownloadPDF">
                            Save PDF
                        </button>
                        <button class="button-danger" @click="clearSheet">
                            Clear Sheet
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
                                        <td>Total Load</td>
                                        <td><input :value="totalELoad" type="number" disabled /></td>
                                    </tr>
                                    <tr>
                                        <td>AD</td>
                                        <td><input :value="borrowTotal" type="number" disabled /></td>
                                    </tr>
                                    <tr>
                                        <td>Total</td>
                                        <td><input :value="totalELoad - borrowTotal" type="number" disabled /></td>
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
                                        <th>266001445</th>
                                        <th>37300247</th>
                                        <th>257283991</th>
                                    </tr>

                                    <tr>
                                        <td>Rs</td>
                                        <td><input v-model="sheet.accountBalance265999891" type="number" /></td>
                                        <td><input v-model="sheet.accountBalance266001445" type="number" /></td>
                                        <td><input v-model="sheet.accountBalance37300247" type="number" /></td>
                                        <td><input v-model="sheet.accountBalance257283991" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Deposit</td>
                                        <td><input v-model="sheet.deposit265999891" type="number" /></td>
                                        <td><input v-model="sheet.deposit266001445" type="number" /></td>
                                        <td><input v-model="sheet.deposit37300247" type="number" /></td>
                                        <td><input v-model="sheet.deposit257283991" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Total Rs</td>
                                        <td><input :value="n(sheet.accountBalance265999891) + n(sheet.deposit265999891)"
                                                type="number" disabled /></td>
                                        <td><input :value="n(sheet.accountBalance266001445) + n(sheet.deposit266001445)"
                                                type="number" disabled /></td>
                                        <td><input :value="n(sheet.accountBalance37300247) + n(sheet.deposit37300247)"
                                                type="number" disabled /></td>
                                        <td><input :value="n(sheet.accountBalance257283991) + n(sheet.deposit257283991)"
                                                type="number" disabled /></td>
                                    </tr>

                                    <tr>
                                        <td>Withdrawl</td>
                                        <td><input v-model="sheet.withdrawl265999891" type="number" /></td>
                                        <td><input v-model="sheet.withdrawl266001445" type="number" /></td>
                                        <td><input v-model="sheet.withdrawl37300247" type="number" /></td>
                                        <td><input v-model="sheet.withdrawl257283991" type="number" /></td>
                                    </tr>

                                    <tr>
                                        <td>Remaining Balance</td>
                                        <td><input
                                                :value="(n(sheet.accountBalance265999891) + n(sheet.deposit265999891)) - n(sheet.withdrawl265999891)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="(n(sheet.accountBalance266001445) + n(sheet.deposit266001445)) - n(sheet.withdrawl266001445)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="(n(sheet.accountBalance37300247) + n(sheet.deposit37300247)) - n(sheet.withdrawl37300247)"
                                                type="number" disabled /></td>
                                        <td><input
                                                :value="(n(sheet.accountBalance257283991) + n(sheet.deposit257283991)) - n(sheet.withdrawl257283991)"
                                                type="number" disabled /></td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>

                        <div class="cards mx">
                            <table>
                                <tbody>
                                    <tr>
                                        <th colspan="2">Cards</th>
                                    </tr>
                                    <tr>
                                        <td>Total</td>
                                        <td><input v-model="sheet.totalCards" type="number" /></td>
                                    </tr>
                                    <tr>
                                        <td>Sell</td>
                                        <td><input v-model="sheet.sellCards" type="number" /></td>
                                    </tr>
                                    <tr>
                                        <td>Remaining</td>
                                        <td><input :value="n(sheet.totalCards) - n(sheet.sellCards)" type="number"
                                                disabled />
                                        </td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>

                <div class="borrow-recovery-container">
                    <div class="borrow-container">
                        <table>
                            <tbody>
                                <tr>
                                    <th colspan="2">Borrow</th>
                                </tr>

                                <tr v-for="(row, index) in sheet.borrow" :key="index">
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
                                        <input type="number" :value="borrowTotal" disabled />
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>

                    <div class="recovery-container">
                        <table>
                            <tbody>
                                <tr>
                                    <th colspan="2">Recovery</th>
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
                                    <input :value="sheet.lastBalances.omni" type="number" disabled />
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
                                    <input :value="sheet.lastBalances.easypaisa" type="number" disabled />
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
                                    <input :value="sheet.lastBalances.jazzcash" type="number" disabled />
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
                                    <input :value="sheet.lastBalances.epaccount" type="number" disabled />
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
                                    <input :value="sheet.lastBalances.jcaccount" type="number" disabled />
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

                <div class="purchasing">
                    <table>
                        <tbody>
                            <tr>
                                <th colspan="2">Purchasing</th>
                            </tr>
                            <tr>
                                <td class="text-sm">UBL Omni Rec</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.omni)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td class="text-sm">EasyPaisa Rec</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.easypaisa)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td class="text-sm">JazzCash Rec</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.jazzcash)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td class="text-sm">EP AC 0333 Rec</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.epaccount)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td class="text-sm">JC Merchant Account</td>
                                <td>
                                    <input :value="getTotalReceiving(sheet.jcaccount)" type="number" disabled />
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
                                <td>Manual Purchasing Total</td>
                                <td>
                                    <input :value="manualPurchasingTotal" type="number" disabled />
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
                                <td>Total Amount</td>
                                <td>
                                    <input :value="n(sheet.previousCash) + cashInfoTotal" type="number" disabled />
                                </td>
                            </tr>

                            <tr>
                                <td>Recovery &amp; Purchasing</td>
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
                                <td>Difference</td>
                                <td>
                                    <input
                                        :value="((n(sheet.previousCash) + cashInfoTotal) - purchasingTotal) - cashTotal"
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
                                <th colspan="2">RED Book</th>
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
                                <td>UBL Omni</td>
                                <td>
                                    <input :value="getTotalSending(sheet.omni)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>EasyPaisa</td>
                                <td>
                                    <input :value="getTotalSending(sheet.easypaisa)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>JazzCash</td>
                                <td>
                                    <input :value="getTotalSending(sheet.jazzcash)" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>EasyPaisa &amp; JazzCash Accounts</td>
                                <td>
                                    <input :value="getTotalSending(sheet.epaccount) +
                                        getTotalSending(sheet.jcaccount)
                                        " type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Recovery</td>
                                <td>
                                    <input :value="recoveryTotal" type="number" disabled />
                                </td>
                            </tr>
                            <tr>
                                <td>Card</td>
                                <td>
                                    <input :value="100 * n(sheet.sellCards)" type="number" disabled />
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
</template>
