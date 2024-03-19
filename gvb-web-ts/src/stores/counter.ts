import {defineStore} from 'pinia'

export interface userInfoType {
    nick_name: string
    role: number  // 角色
    user_id: number  // 用户id
    avatar: string
    token: string
}

const theme: boolean = true
const collapsed: boolean = false
const userInfo: userInfoType = {
    nick_name: "plktime1043",
    role: 0,
    user_id: 0,
    avatar: "http://localhost:5173/image/user_head.jpg",
    token: "",
}

export const useStore = defineStore('counter', {
    state() {
        return {
            theme: theme,
            collapsed: collapsed, // 侧边栏默认一开始不折叠
            userInfo: userInfo,
        }
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
            // console.log(obj)
        },
        setCollapsed(collapsed: boolean) {
            this.collapsed = collapsed
        }
    },
    getters: {
        themeString(): string {
            return this.theme ? "light" : "dark"
        }
    }
})
