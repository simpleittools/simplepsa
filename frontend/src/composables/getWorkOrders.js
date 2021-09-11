import {ref} from "vue";


const getWorkOrders = () => {
    const workOrders = ref([])
    const error = ref(null)

    const load = () => {
        fetch("http://localhost:3330/api/workorder")
            .then(response => response.json())
            .then(data => (workOrders.value = data))
    }

    return {
        workOrders,
        error,
        load
    }
}
export default getWorkOrders