<template>
  <div class="ques-list">
    <el-dialog
      v-model="sortDialog"
      :title="selectSort.id ? 'Modify Category' : 'Add Category'"
      width="30%"
      :before-close="closeSort"
    >
      <el-input
        v-model="selectSort.name"
        placeholder="Please enter the category name"
      ></el-input>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="closeSort">Cancel</el-button>
          <el-button type="primary" @click="subSort">Done</el-button>
        </span>
      </template>
    </el-dialog>
      <el-dialog
      v-model="quesDialog"
      :title="selectQues.id ? 'Modify Problem' : 'Add Problem'"
      width="70%"
      :close-on-click-modal="false"
      :before-close="closeQues"
    >
      <addQues @cancel="closeQues" :selectQues="selectQues" @submit="subQues" :sortList="sortList" v-if="quesDialog"></addQues> 
      
      
    </el-dialog>
    <div class="flex-box">
      <div class="sort-list">
        <h3>
          Category<el-icon @click="sortDialog = true"><circle-plus /></el-icon>
        </h3>

        <div @click="getProblem(null)" :class="!actSort ? 'act' : ''">All</div>
        <div
          v-for="item in sortList"
          :key="item.id"
          @click="getProblem(item.identity)"
          :class="actSort == item.identity ? 'act' : ''"
        >
          <span>{{ item.name }}</span>

          <div class="edit">
            <el-icon title="Delete" @click.stop="delSort(item)"
              ><delete
            /></el-icon>
            <el-icon title="Modify" @click.stop="editSort(item)"
              ><edit
            /></el-icon>
          </div>
        </div>
      </div>
      <div class="list" v-loading="loading">
        <h3>
          <el-input
            style="width: 200px; margin-right: 10px"
            v-model="keyword"
            clearable
            @change="getProblem(actSort)"
            @clear="getProblem(null)"
            placeholder="Search"
            :suffix-icon="Search"
          />
          <el-button type="primary" @click="quesDialog=true">Add Problem</el-button>
        </h3>
        <div class="scroll">
          <div class="item" v-for="item in quesList" :key="item.id">
            <div>
              <div class="title">
                <b>{{ item.title }}</b>
                <div class="sort">
                  <span v-for="mi in item.problem_categories" :key="mi.id" >
                    <b v-if="mi.category_basic">{{ mi.category_basic.name }}</b>
                  </span>
                </div>
              </div>
              <div class="msg">
                <span>Create At: {{ item.created_at }}</span>
                <span>Submitions Num: {{ item.submit_num }}</span>
                <span>Pass Num: {{ item.pass_num }}</span>
              </div>
            </div>
            <div class="edit">
              <el-icon title="Delete" @click="delQues(item)"><delete /></el-icon>
              <el-icon title="Modify" @click="editQues(item)"><edit /></el-icon>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="pagi">
      <el-pagination
        v-model:currentPage="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        layout="total,sizes, prev, pager, next"
        :total="total"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>
<script lang="ts" setup>
import { reactive, ref } from "@vue/reactivity";
import { Edit, Delete, List, CirclePlus } from "@element-plus/icons-vue";
import { ElMessageBox, ElMessage } from "element-plus";
import api from "../../api/api.js";
import { useRouter } from "vue-router";
import addQues from './add.vue'
import modQues from './modefy.vue'
import { registerLoading } from "echarts";
const router = useRouter();
const loading = ref(false);
const quesList = ref([]);
const sortList = ref([]);
const sortDialog = ref(false);
const selectSort = ref({ name: "",id:'' });

const quesDialog=ref(false)
const selectQues=ref({})

const pageSize = ref(10);
const currentPage = ref(1);
const total = ref(0);
const actSort = ref<null | number>(null);
const keyword = ref("");
const closeQues=()=>{
  quesDialog.value=false
  selectQues.value={}
}
const subQues=()=>{
closeQues()
getProblem(actSort.value)
}
const editQues=(item:any)=>{
   
    api.getProblemDetail({
        identity:item.identity
    }).then(res=>{
            if(res.data.data){
                selectQues.value=res.data.data
                quesDialog.value=true
            }
    })
 
// selectQues.value=item

}
const handleSizeChange = (val: number) => {
  getProblem(actSort.value);
};
const editSort = (item: any) => {
  selectSort.value = {...item};
  sortDialog.value = true;
};
const handleCurrentChange = (val: number) => {
  getProblem(actSort.value);
};
const closeSort = () => {
  sortDialog.value = false;
  selectSort.value = { name: "",id:'' };
};
const delSort = (item: any) => {
  ElMessageBox.confirm("Are you sure you want to delete this category?", "Hint", {
    confirmButtonText: "Yes",
    cancelButtonText: "No",
    type: "warning",
  })
    .then(() => {
      api.delSort({ identity: item.identity }).then((res) => {
        if (res.data.code == 200) {
          ElMessage.success(res.data.msg);
          getSortList();
        } else {
          ElMessage.warning(res.data.msg);
        }
      });
    })
    .catch(() => {
      // catch error
    });
};
const delQues = (item: any) => {
  ElMessageBox.confirm("Are you sure you want to delete this problem?", "Hint", {
    confirmButtonText: "Yes",
    cancelButtonText: "No",
    type: "warning",
  })
    .then(() => {
      api.delProblem({ identity: item.identity }).then((res) => {
        if (res.data.code == 200) {
          ElMessage.success(res.data.msg);
          getProblemList();
         
        } else {
          ElMessage.warning(res.data.msg);
        }
      });
    })
    .catch(() => {
      // catch error
    });
};
const getProblem = (sortId: number | null) => {
  
  actSort.value = sortId;
  loading.value = true;
  api
    .getProblemList({
      category_identity: sortId,
      size: pageSize.value,
      page: currentPage.value,
      keyword: keyword.value,
    })
    .then((res) => {
      loading.value = false;
      if (res && res.data) {
        quesList.value = res.data.data.list;
        total.value = res.data.data.count;
      }
      console.log(res);
    });
};
getProblem(null);
const getSortList = () => {
  api.getSortList({}).then((res) => {
    if (res && res.data) {
      sortList.value = res.data.data.list;
    }
  });
};
getSortList();

const getProblemList = () => {
  api.getProblemList({}).then((res) => {
    if (res && res.data) {
      quesList.value = res.data.data.list;
    }
  });
};
getProblemList();

const subSort = () => {
  if (selectSort.value.name) {
    if (selectSort.value.id) {
      api.editSort(selectSort.value).then((res) => {
        if (res.data.code == 200) {
          ElMessage.success(res.data.msg);
          closeSort();
          getSortList()
        } else {
          ElMessage.warning(res.data.msg);
        }
      });
    } else {
      api.addSort({ name: selectSort.value.name }).then((res) => {
        if (res.data.code == 200) {
          ElMessage.success(res.data.msg);
          closeSort();
          getSortList()

        } else {
          ElMessage.warning(res.data.msg);
        }
      });
    }
  } else {
    ElMessage.warning("Please enter category name");
  }
};
const toDetail = (item: any) => {
  router.push({
    path: "/questionDetail",
    query: item,
  });
};
</script>
<style scoped lang="scss">
.scroll {
  height: calc(100% - 60px);
  overflow: auto;
}
.ques-list {
  height: 100%;
  display: flex;
  // overflow: hidden;
  justify-content: space-between;
  flex-direction: column;
}
.flex-box {
  display: flex;
  height: calc(100% - 60px);
}
.list {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100%;
  h3 {
    // background-color: #0087bf;
    // color: white;
    // border-bottom: 1px solid #0087bf;

    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
  }
}
.pagi {
  text-align: center;
  display: flex;
  justify-content: center;
  padding: 10px 0;
  border-top: 1px solid #eee;
}
.item {
  padding: 20px;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .edit {
    width: 80px;
    display: flex;
    justify-content: space-between;
    color: #13b6f9;
    cursor: pointer;
  }
  .title {
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    font-size: 18px;
    b {
      cursor: pointer;
    }
    span {
      background-color: #d3ebff;
      font-size: 12px;
      margin: 0 10px;
      padding: 4px;
      border-radius: 5px;
      border: 1px solid #62a9ff;
      color: #62a9ff;
    }
  }
  .msg {
    
    font-size: 14px;
    color: #999;
    span {
      margin-right: 20px;
    }
  }
}
.sort-list {
  font-size: 20px;
  // display: flex;
  width: 250px;
  align-items: center;
  justify-content: space-between;
  //   background-color: #0087bf;
  //   color: white;
  border-right: 1px solid #0087bf;

  // padding: 20px;
  & > div {
    display: flex;
    justify-content: space-between;
    padding: 10px;
  }
  .edit {
    font-size: 20px;
    width: 20%;
    display: flex;
    justify-content: space-between;
  }
  h3 {
    // background-color: #0087bf;
    border-bottom: 1px solid #0087bf;
    // color: white;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 14px;
  }
  div {
    color: #13b6f9;

    cursor: pointer;
    &.act {
      // color: black;
      background-color: #c4e9f9;
    }
  }
}
</style>