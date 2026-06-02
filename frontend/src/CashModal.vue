<script setup lang="ts">
import { reactive, computed } from "vue";

defineProps<{
    modelValue: boolean;
}>();

const emit = defineEmits<{
    (e: "update:modelValue", value: boolean): void;
}>();

const amounts = reactive<(number | null)[]>(Array(8).fill(null));

const cash = reactive({
    cash5000: null as number | null,
    cash1000: null as number | null,
    cash500: null as number | null,
    cash100: null as number | null,
    cash50: null as number | null,
    cash20: null as number | null,
    cash10: null as number | null,
    cash5: null as number | null,
});

const n = (value: number | null | undefined): number => {
    return Number(value) || 0;
};

const totalAmount = computed((): number => {
    return amounts.reduce<number>((sum, value) => {
        return sum + n(value);
    }, 0);
});

const cashReceived = computed((): number => {
    return (
        n(cash.cash5000) * 5000 +
        n(cash.cash1000) * 1000 +
        n(cash.cash500) * 500 +
        n(cash.cash100) * 100 +
        n(cash.cash50) * 50 +
        n(cash.cash20) * 20 +
        n(cash.cash10) * 10 +
        n(cash.cash5) * 5
    );
});

const changeAmount = computed((): number => {
    return cashReceived.value - totalAmount.value;
});

const clearAll = () => {
    amounts.fill(null);
    Object.keys(cash).forEach((key) => {
        (cash as any)[key] = null;
    });
};

const closeModal = () => {
    emit("update:modelValue", false);
};
</script>

<template>
    <div v-if="modelValue" class="modal-overlay">
        <div class="modal-box">
            <div class="modal-header">
                <h2>Cash Calculator</h2>
                <button class="close-btn" @click="closeModal">✕</button>
            </div>
            <div class="content">
                <div class="card">
                    <h3>Total Amount</h3>
                    <table class="table">
                        <tbody>
                            <tr v-for="(amount, index) in amounts" :key="index">
                                <td>
                                    <input v-model.number="amounts[index]" type="number" min="0" />
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <button class="clear-btn" @click="clearAll">Clear All</button>
                </div>

                <div class="card">
                    <h3>Cash Details</h3>
                    <table class="table">
                        <thead>
                            <tr>
                                <th>Note</th>
                                <th>Qty</th>
                                <th>Total</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="note in [5000, 1000, 500, 100, 50, 20, 10, 5]" :key="note">
                                <td>{{ note }}</td>
                                <td>
                                    <input v-model.number="cash[`cash${note}` as keyof typeof cash]" type="number"
                                        min="0" />
                                </td>
                                <td style="text-align: center;">
                                    {{ (n(cash[`cash${note}` as keyof typeof cash]) * note).toLocaleString() }}
                                </td>
                            </tr>
                        </tbody>
                    </table>

                    <div class="summary">
                        <div class="summary-row">
                            <span>Total Amount</span>
                            <strong>{{ totalAmount.toLocaleString() }}</strong>
                        </div>
                        <div class="summary-row">
                            <span>Cash Received</span>
                            <strong>{{ cashReceived.toLocaleString() }}</strong>
                        </div>
                        <div class="summary-row">
                            <span>Change</span>
                            <strong :class="{
                                positive: changeAmount >= 0,
                                negative: changeAmount < 0
                            }">
                                {{ changeAmount.toLocaleString() }}
                            </strong>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, .55);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999;
}

.modal-box {
    width: 1000px;
    max-width: 95vw;
    background: white;
    border-radius: 12px;
    padding: 20px;
}

.modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}

.modal-header h2 {
    margin: 0;
}

.close-btn {
    width: 35px;
    height: 35px;
    border: none;
    background: #ef4444;
    color: white;
    border-radius: 6px;
    cursor: pointer;
    font-size: 12px;
}

.content {
    display: grid;
    grid-template-columns: 320px 1fr;
    gap: 20px;
}

.card {
    border: 1px solid #ddd;
    border-radius: 10px;
    padding: 15px;
    background: white;
}

.card h3 {
    margin-top: 0;
    margin-bottom: 15px;
}

.table {
    width: 100%;
    border-collapse: collapse;
}

.table th,
.table td {
    border: 1px solid #ddd;
    padding: 8px;
}

.table th {
    background: #f5f5f5;
}

.table input {
    width: 100%;
    box-sizing: border-box;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 0.8rem;
}

.footer-row {
    background: #fafafa;
}

.clear-btn {
    width: 100%;
    margin-top: 15px;
    padding: 12px;
    border: none;
    background: #dc2626;
    color: white;
    border-radius: 6px;
    cursor: pointer;
    font-weight: 600;
    font-size: 0.8rem;
}

.clear-btn:hover {
    background: #b91c1c;
}

.summary {
    margin-top: 20px;
    display: flex;
    flex-direction: column;
    gap: 10px;
}

.summary-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px;
    background: #f7f7f7;
    border-radius: 6px;
}

.positive {
    color: #16a34a;
}

.negative {
    color: #dc2626;
}

@media (max-width: 768px) {
    .content {
        grid-template-columns: 1fr;
    }
}
</style>