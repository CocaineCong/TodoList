<template>
    <div class="login_from">
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="账号" prop="user">
                <el-input v-model="ruleForm.user"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="pass">
                <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
                <el-button @click="toRegister">没有账号？去注册</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        data() {
            var checkAge = (rule, value, callback) => {
                if (!value) {
                    return callback(new Error('账号不能为空'));
                }
            };
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
            return {
                ruleForm: {
                    pass: '',
                    user: ''
                },
                rules: {
                    pass: [
                        {validator: validatePass, trigger: 'blur'}
                    ],
                    user: [
                        {validator: checkAge, trigger: 'blur'}
                    ]
                }
            };
        },
        methods: {
            submitForm(formName) {
                console.log("asd")

                this.$axios.post('/api/v1/user/login', {
                    user_name: this.ruleForm.user,
                    password: this.ruleForm.pass
                }).then(res => {
                    this.$message({
                        message: '登陆成功！',
                        type: 'success'
                    });
                    console.log(res)
                    sessionStorage.setItem("token", res.token)
                    sessionStorage.setItem('username',res.user.user_name)
                    //跳转home页面
                    this.$router.push("/home")
                }).catch(error =>{
                    this.$message({
                        message:'登陆失败！',
                        type: 'error'
                    })
                    console.log(error)
                })


                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        // alert('submit!');

                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            },
            toRegister() {
                this.$router.push('register')
            }
        }
    }
</script>

<style scoped>
.login_from {
    width: 500px;
    height: 350px;
    display: flex;

    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 30px;
    border: 2px solid sandybrown;
    background-color: whitesmoke;


}
    .demo-ruleForm{
        display: flex;
        flex-direction: column;
        justify-content: center;
    }
</style>

