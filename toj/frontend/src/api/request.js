/****   request.js   ****/
// 导入axios
import axios from 'axios'
import qs from 'qs'
import { ElMessage } from 'element-plus'
// import store from "../store";
// import CryptoJs from 'crypto-js'
// use element-ui Message Remainder


// var msk=document.getElementsByClassName('axios-mask')[0]
const service = axios.create({
 
	baseURL: 'http://124.221.199.57:8019',
	// timeout 3s，
	timeout: 300 * 1000
})
// 2.request interceptor
service.interceptors.request.use(config => {
		if (config.method === 'get') {
			config.paramsSerializer = function(params) {
				return qs.stringify(params, {
					arrayFormat: 'comma'
				})
			}
		}
		const token=localStorage.token
		if(token){
			
			config.headers['Authorization']=token
		}
	return config
}, error => {
	Promise.reject(error)
})

// response interceptor
service.interceptors.response.use((config) => {
	return config
}, (error) => {
	 
	if (error.response) {
		const errorMessage = error.response.data === null ? 'The system is abnormal, please contact the site administrator!' : error.response.data.message
		switch (error.response.status) {
			case 404:
				ElMessage('Sorry, resources not found')

				break
			case 403:
				ElMessage('Sorry, you do not have this operation right now')


				break
			case 401:
				ElMessage('Sorry, the authentication has expired, please log in again')
				// let aid=getQueryVariable('corpId')
				// localStorage.removeItem(aid)
				// alert(aid)

				break
			default:
				if (errorMessage === 'refresh token not found') {
					ElMessage('Login has expired, please login again')
						 

				} else {
					ElMessage(errorMessage)


				}
				break
		}
	}
	return Promise.reject(error)
})

export default service
