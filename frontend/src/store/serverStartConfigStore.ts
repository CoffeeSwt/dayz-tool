import { defineStore } from "pinia"
import { reactive } from "vue"
import { GetServerStartConfig } from '../../wailsjs/go/service/ConfigService'

export const useServerStartConfigStore = defineStore('serverStartConfig', () => {
    const config = reactive({
        dayzserverpath: '',
        clientmods: '',
        servermods: '',
        port: ''
    })

    const init = async () => {
        const res = await GetServerStartConfig()
        console.log(res, config)
    }
    return { config, init }
})