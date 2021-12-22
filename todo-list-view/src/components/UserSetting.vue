<template>
  <div>
  <el-header style="text-align: right; font-size: 20px">
    <el-dropdown>
      <i class="el-icon-s-custom" style="margin-right: 30px"></i>
      <el-dropdown-menu class="el-aside" slot="dropdown">
<!--        <el-dropdown-item>我的</el-dropdown-item>-->
        <el-dropdown-item @click.native="viewDialogUpdate">修改密码</el-dropdown-item>
        <el-dropdown-item @click.native="logout">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
    <span>{{ username}}</span>
  </el-header>

    <el-dialog
        :visible.sync="dialogUpdatePass"
        width="25%"
        :close-on-click-modal="true"
        :show-close="true"
        :before-close="handleClose">
      <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
        <el-form-item label="密码" prop="pass">
          <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="checkPass">
          <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
          <el-button @click="resetForm('ruleForm')">重置</el-button>
        </el-form-item>
      </el-form>

    </el-dialog>

  </div>
</template>


<script>
export default {
  name: "UserSetting",
  data(){
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.pass) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    return{
      username : sessionStorage.getItem('username'),
      //修改的对话框
      dialogUpdatePass: false,
      ruleForm: {
        pass: '',
        checkPass: '',
      },
      rules: {
        pass: [
          { validator: validatePass, trigger: 'blur' }
        ],
        checkPass: [
          { validator: validatePass2, trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          //调用修改密码接口
          this.clickUpdatePassAPI();
          this.dialogUpdatePass= false;
          //跳转home页面
          this.logout()
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    //展示更新对话框
    viewDialogUpdate(){
      this.dialogUpdatePass= true
    },
    //更新密码API
    clickUpdatePassAPI(){
      this.$axios.post('/api/v1/user/updatepass', {
        user_name: this.username,
        password: this.ruleForm.pass
      }).then(res => {
        console.log(res)
        this.$message({
          message: '修改成功！',
          type: 'success'
        });

      }).catch(error =>{
        this.$message({
          message:'修改失败！',
          type: 'error'
        })
        console.log(error)
      })
    },
    //右上角关闭按钮
    handleClose(done) {
      this.$confirm('确认关闭？')
          .then(() => {
            done();
          })
          .catch(() => {});
    },
    logout(){
      //删除用户信息
      sessionStorage.removeItem("token")
      sessionStorage.removeItem("name")
      this.$message({
        message: '注销成功！',
        type: 'success'
      });
      //返回登录页面
      this.$router.push("/login")
    }
  },
  created() {
    console.log(this.$parent);
  }
}
</script>

<style scoped>

</style>