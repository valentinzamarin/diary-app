import { reactive } from 'vue'

const alerts = reactive([])

export function useAlert() {
    const showAlert = (message, type = 'info') => {
        const id = Date.now()
        alerts.push({ id, message, type })
        setTimeout(() => {
            alerts.splice(alerts.findIndex(a => a.id === id), 1)
        }, 3000)
    }
    return { showAlert, alerts }
}