import axios from "axios"
import { Message } from "element-ui"
const baseUrl = "http://127.0.0.1:3000/"
const instance = axios.create({
    baseURL: process.env.NODE_ENV == 'development' ? '' : baseUrl,
    timeout: 6000,
});

//请求拦截，添加token
instance.interceptors.request.use(config => {
    const token = sessionStorage.getItem('token')
    if(token) {
        config.headers['authorization'] = token
    }
    return config
})

//响应拦截，处理请求结果
instance.interceptors.response.use(res => {
    return new Promise((resolve) => {
        if(res.data.status == 200) {
            resolve(res.data.data)
        }else {
            console.log(res)
            Message.closeAll()
            Message({
                message: res.data.msg,
                type: 'error'
            })
            // reject(res.data.msg)
        }
    })
})
export default instance;