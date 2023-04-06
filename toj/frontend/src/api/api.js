import http from './http.js'

 
export default{
	//  get request example
	getProblemList(param){//Problem List
		return http.get(`/problem-list`,param)
	},
	getProblemDetail(param){//Problem Detail
		return http.get(`/problem-detail`,param)
	},
	getSortList(param){//Category
		return http.get(`/category-list`,param)
	},
	getRankList(param){//Rank List
		return http.get(`/rank-list`,param)
	},
	getSubmitList(param){//Submit List
		return http.get(`/submit-list`,param)
	},
	sendCode(param){//Send Code
		return http.postUncode(`/send-code`,param)
	},
	login(param){//login
		return http.postUncode(`/login`,param)
	},
	register(param){//register
		return http.postUncode(`/register`,param)
	},
	forgetPassword(param){//register
		return http.postUncode(`/forget-password`,param)
	},
	delSort(param){//category delete 
		return http.delete(`/admin/category-delete`,param)
	},
	addSort(param){//category create
		return http.postUncode(`/admin/category-create`,param)
	},
	addProblem(param){//problem create
		return http.post(`/admin/problem-create`,param)
	},
	editProblem(param){//problem modify
		return http.putJson(`/admin/problem-modify`,param)
	},
	editSort(param){//category modify
		return http.put(`/admin/category-modify`,param)
	},
	submitCode(param,id){//submit code
		return http.postJson(`/user/submit?problem_identity=${id}`,param)
	},
	// file upload
	uploadFile(param){
		return http.upFile('',param)
	},
	
	
 
	 
}
