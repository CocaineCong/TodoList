<template>
  <div>
    <el-container style="height: 100%; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu :default-openeds="['1', '3']">
          <el-submenu index="1">
            <template slot="title"><i class="el-icon-tickets">作业管理</i></template>
            <el-menu-item-group>
              <template slot="title">备忘录</template>
              <el-menu-item index="1-1">未完成列表</el-menu-item>
              <el-menu-item index="1-2">以完成列表</el-menu-item>
            </el-menu-item-group>
          </el-submenu>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header style="text-align: right; font-size: 20px">
        <UserSetting></UserSetting>
        </el-header>
        <el-main class="internal-content">
          <el-table :data="tableData">
            <el-table-column prop="start_time" label="日期" width="100%">
            </el-table-column>
            <el-table-column label="姓名" width="100%">
              {{ username }}
            </el-table-column>
            <el-table-column prop="title" label="标题">
            </el-table-column>
            <el-table-column prop="content" label="备忘录简述">
            </el-table-column>
            <el-table-column>
              <template slot-scope="scope">
                <el-tooltip placement="top"  effect="light">
                  <div slot="content">编辑</div>
                  <el-button @click="clickEditTable(scope.row)" type="primary" icon="el-icon-edit" circle></el-button>
                </el-tooltip>
                <el-tooltip placement="top" effect="light">
                  <div slot="content">详情</div>
                  <el-button @click="clickViewInfo(scope.row.content)" type="info" icon="el-icon-info" circle></el-button>
                </el-tooltip>
                <el-tooltip placement="top" effect="light">
                  <div slot="content">删除</div>
                  <el-button type="danger" @click="clickDeleteTable(scope.row)" icon="el-icon-delete" circle></el-button>
                </el-tooltip>

              </template>
            </el-table-column>
          </el-table>
        </el-main>
      </el-container>
    </el-container>

    <el-dialog title="编辑备忘录" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">
      <div class="el-input-group">
        标题:
        <el-input v-model="diglogData.title"></el-input>
      </div>
      <div class="">
        文本：
        <el-input
            type="textarea"
            :rows="2"
            placeholder="请输入内容"
            v-model="diglogData.content">
        </el-input>
      </div>
      <span slot="footer" class="dialog-footer">
                   <el-button @click="dialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="clickEdit">确 定</el-button>
                  </span>
    </el-dialog>


    <el-dialog
        title="提示"
        :visible.sync="dialogVisibleDelete"
        width="30%"
        :before-close="handleClose">
      <span>删除后数据无法恢复，您确认删除吗？</span>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleDelete = false">取 消</el-button>
        <el-button type="primary" @click="clickDelete">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>


<script lang="js">
import UserSetting from "./UserSetting";
export default {
  name: "home",
  components: {UserSetting},
  data() {

    return {
      //修改的对话框
      dialogVisible: false,
      //删除按钮的确认对话框
      dialogVisibleDelete: false,
      content: '',
      //收藏用户名
      username: sessionStorage.getItem('username'),
      //后端返回的item
      tableData: [],
      //修改填写信息
      textarea: '',
      //当前选中的信息
      diglogData: {}
    }
  },
  methods: {
    onCreate() {
      this.$axios.get('/api/v1/tasks', {}).then(res => {
        this.total = res.total
        this.tableData = res.item
        // item": [
        // {
        //   "id": 1,
        //     "title": "注意上课时间",
        //     "content": "明天上午第二节有课",
        //     "view": 0,
        //     "status": 0,
        //     "created_at": 1638198036,
        //     "start_time": 1638198035,
        //     "end_time": 0
        // },]
      }).catch(error => {
        this.$message({
          message: '查询数据失败！',
          type: 'error'
        })
        console.log(error)
      })
    },
    clickEdit(id,title,content) {
      console.log(id, title, content)
      this.dialogVisible = false;
      this.$axios.put('api/v1/task/' + this.diglogData.id, {
        title : this.diglogData.title,
        content: this.diglogData.content
      }).then(res => {
        console.log(res)
        this.$message({
          message: '修改数据成功!',
          type: 'success'
        })
        console.log(this.content)
      }).catch(error => {
        this.$message({
          message: '修改数据失败！',
          type: 'error'
        })
        console.log(error)
      })


    },
    handleClose(done) {
      this.$confirm('确认关闭？')
          .then(() => {
            done();
          })
          .catch(() => {});
    },
    clickDelete() {
      this.dialogVisibleDelete = false
      this.$axios.delete('api/v1/task/' + this.diglogData.id, {

      }).then(res => {
        console.log(res)
        this.$message({
          message: '删除数据成功!',
          type: 'success'
        })
        console.log(this.content)
        //删除成功之后重新查询一下
        this.onCreate()
      }).catch(error => {
        this.$message({
          message: '删除数据失败！',
          type: 'error'
        })
        console.log(error)
      })
    },
    clickViewInfo(content) {
      this.$alert(content, '备忘录详细信息', {
        confirmButtonText: '确定',
        // callback: action => {
        //   this.$message({
        //     type: 'info',
        //     message: `action: ${ action }`
        //   });
        // }
      });
    },

    //点击展示修改diglog
    clickEditTable(data) {
      this.diglogData = data
      this.dialogVisible = true
    },
    //点击展示删除diglog
    clickDeleteTable(data) {
      this.diglogData = data
      this.dialogVisibleDelete = true
    }
  },
  created() {
    this.onCreate()
    // this.$router.push('/home/news');
  },
};
</script>


<style>
.internal-content {
  display: flex;
  width: 1600px;
}

.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}

.el-aside {
  color: #333;
}
</style>
