<script setup lang="ts">
import { reactive, computed } from "vue";

defineProps<{
    modelValue: boolean;
}>();

const emit = defineEmits([
    "update:modelValue",
]);

const modalSheet = reactive({
    modalCash5000: null as number | null,
    modalCash1000: null as number | null,
    modalCash500: null as number | null,
    modalCash100: null as number | null,
    modalCash50: null as number | null,
    modalCash10: null as number | null,

    modalTotal: null as number | null,
});

const n = (v: number | null | undefined) => {
    return isNaN(Number(v)) ? 0 : Number(v);
};

const receiveTotal = computed(() => {
    return (
        5000 * n(modalSheet.modalCash5000) +
        1000 * n(modalSheet.modalCash1000) +
        500 * n(modalSheet.modalCash500) +
        100 * n(modalSheet.modalCash100) +
        50 * n(modalSheet.modalCash50) +
        10 * n(modalSheet.modalCash10)
    );
});

const baqaya = computed(() => {
    return receiveTotal.value - n(modalSheet.modalTotal);
});

const closeModal = () => {
    emit("update:modelValue", false);
};
</script>

<template>
    <div v-if="modelValue" class="modal-overlay">
        <div class="modal-box">
            <div class="modal-header">
                <h2>Quick Cash Calculator</h2>
                <button class="close-btn" @click="closeModal">
                    ✕
                </button>
            </div>
            <table class="table">
                <tbody>
                    <tr>
                        <th colspan="4">
                            Cash Detail
                        </th>
                    </tr>
                    <tr>
                        <td>5000</td>
                        <td>x</td>
                        <td>
                            <input v-model.number="modalSheet.modalCash5000" type="number" min="0" />
                        </td>
                        <td>
                            <input :value="5000 * n(modalSheet.modalCash5000)" disabled />
                        </td>
                    </tr>
                    <tr>
                        <td>1000</td>
                        <td>x</td>
                        <td>
                            <input v-model.number="modalSheet.modalCash1000" type="number" />
                        </td>
                        <td>
                            <input :value="1000 * n(modalSheet.modalCash1000)" disabled />
                        </td>
                    </tr>

                    <tr>
                        <td>500</td>
                        <td>x</td>
                        <td>
                            <input v-model.number="modalSheet.modalCash500" type="number" />
                        </td>
                        <td>
                            <input :value="500 * n(modalSheet.modalCash500)" disabled />
                        </td>
                    </tr>
                    <tr>
                        <td>100</td>
                        <td>x</td>
                        <td>
                            <input v-model.number="modalSheet.modalCash100" type="number" />
                        </td>
                        <td>
                            <input :value="100 * n(modalSheet.modalCash100)" disabled />
                        </td>
                    </tr>

                    <tr>
                        <td>50</td>
                        <td>x</td>
                        <td>
                            <input v-model.number="modalSheet.modalCash50" type="number" />
                        </td>
                        <td>
                            <input :value="50 * n(modalSheet.modalCash50)" disabled />
                        </td>
                    </tr>
                    <tr>
                        <td>10</td>
                        <td>x</td>
                        <td>
                            <input v-model.number="modalSheet.modalCash10" type="number" />
                        </td>
                        <td>
                            <input :value="10 * n(modalSheet.modalCash10)" disabled />
                        </td>
                    </tr>
                    <tr>
                        <td colspan="3">
                            Total Cash Received
                        </td>
                        <td>
                            <input :value="receiveTotal" disabled />
                        </td>
                    </tr>

                </tbody>
            </table>
            <table class="table mt">
                <tbody>
                    <tr>
                        <td>Receive</td>

                        <td>
                            <input :value="receiveTotal" disabled />
                        </td>
                    </tr>
                    <tr>
                        <td>Total</td>

                        <td>
                            <input v-model.number="modalSheet.modalTotal" type="number" placeholder="Customer Total" />
                        </td>
                    </tr>
                    <tr>
                        <td>Baqaya Cash</td>
                        <td>
                            <input :value="baqaya" disabled />
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<style scoped>
.modal-overlay {
    position: fixed;
    inset: 0;
    background: rgba(0, 0, 0, 0.5);
    font-size: 1rem;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 1rem;
    z-index: 9999;
}

.modal-box {
    width: 700px;
    max-width: 95%;
    background: white;
    font-size: 1rem;
    padding: 20px;
    border-radius: 10px;
}

.modal-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 1rem;
    margin-bottom: 15px;
}

.close-btn {
    width: 30px;
    height: 30px;
    border: none;
    border-radius: 5px;
    background: red;
    color: white;
    cursor: pointer;
}

.table {
    width: 100%;
    border-collapse: collapse;
}

.table td,
.table th {
    border: 1px solid #ccc;
    padding: 8px;
    font-size: 1rem;
}

.table input {
    width: 100%;
    box-sizing: border-box;
    font-size: 1rem;
}

.mt {
    margin-top: 20px;
}
</style>