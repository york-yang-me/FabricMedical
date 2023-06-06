<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="patient" prop="patient">
        <el-select v-model="ruleForm.patient" placeholder="please choose patient" @change="selectGet">
          <el-option
            v-for="item in accountList"
            :key="item.accountId"
            :label="item.userName"
            :value="item.accountId"
          >
            <span style="float: left">{{ item.userName }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.accountId }}</span>
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="total length" prop="totalLength">
        <el-input-number v-model="ruleForm.totalLength" :precision="2" :step="0.1" :min="0" />
      </el-form-item>
      <el-form-item label="dna contents" prop="dnaContents">
        <el-input-number v-model="ruleForm.dnaContents" :precision="2" :step="0.1" :min="0" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">create</el-button>
        <el-button @click="resetForm('ruleForm')">reset</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createRealSequence } from '@/api/realSequence'

export default {
  name: 'AddRealSequence',
  data() {
    var checkLength = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('must be bigger than 0'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        patient: '',
        totalLength: 0,
        dnaContents: 0
      },
      accountList: [],
      rules: {
        patient: [
          { required: true, message: 'please choose patient', trigger: 'change' }
        ],
        totalLength: [
          { validator: checkLength, trigger: 'blur' }
        ],
        dnaContents: [
          { validator: checkLength, trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'accountId'
    ])
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        // filter the admin
        this.accountList = response.filter(item =>
          item.userName !== 'admin'
        )
      }
    })
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('Whether to create immediately?', 'tip', {
            confirmButtonText: 'confirm',
            cancelButtonText: 'cancel',
            type: 'success'
          }).then(() => {
            this.loading = true
            createRealSequence({
              accountId: this.accountId,
              patient: this.ruleForm.patient,
              totalLength: this.ruleForm.totalLength,
              dnaContents: this.ruleForm.dnaContents
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: 'create success!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: 'create failed!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: 'cancelled'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.ruleForm.patient = accountId
    }
  }
}
</script>

<style scoped>
</style>
