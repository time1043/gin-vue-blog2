import {ref, computed} from 'vue'
import {defineStore} from 'pinia'
import {getFiles} from "@arco-design/web-vue/es/upload/utils";

const theme = true

export const useStore = defineStore('counter', {
    state() {
        return {theme: theme}
    },
    actions: {
        setTheme(localTheme?: boolean) {
            if (localTheme !== undefined) {  // 传参了
                this.theme = localTheme
            } else {  // 没传参则取反
                this.theme = !this.theme
            }
            document.documentElement.style.colorScheme = this.themeString
            document.body.setAttribute("arco-theme", this.themeString)

            localStorage.setItem("theme", JSON.stringify(this.theme))
        },
        loadTheme() {
            let val = localStorage.getItem("theme")
            if (val === null) {
                return
            }
            try {
                this.theme = JSON.parse(val)
                this.setTheme(this.theme)
            } catch (e) {
                return;
            }
            let obj = JSON.parse(val)
            console.log(obj)
        }
    },
    getters: {
        themeString(): string {
            return this.theme ? "light" : "dark"
        }
    }
})
