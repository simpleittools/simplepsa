import { ref } from 'vue'

const addClient = (collection) => {

    const error = ref(null)
    const isPending = ref(false)


    // add a new document
    const addClient = async (c) => {
        error.value = null
        isPending.value = true


        try {
            const res = await "http://localhost:3330/api/company/create".add(c)
            isPending.value = false
            return res

        }
        catch(err) {
            console.log(err.message)
            error.value = 'could not send the message'
            isPending.value = false
        }
    }

    return { error, addClient, isPending }

}

export default addClient