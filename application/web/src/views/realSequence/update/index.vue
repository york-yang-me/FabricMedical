<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="owner" prop="owner">
        <el-select v-model="ruleForm.owner" placeholder="please choose owner" @change="selectGet">
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
        <el-input-number v-model="ruleForm.totalLength" :step="1" :min="0"/>
      </el-form-item>
      <el-form-item label="dna contents" prop="dnaContents">
        <el-input type="textarea" placeholder="input dna sequence contents" v-model="ruleForm.dnaContents" />
      </el-form-item>
      <el-form-item label="dna description" prop="description">
        <el-input type="textarea" placeholder="input dna sequence description" v-model="ruleForm.description" />
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
  name: 'UpdateRealSequence',
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
        owner: '',
        totalLength: 0,
        dnaContents: ''
      },
      accountList: [],
      rules: {
        owner: [
          { required: true, message: 'please choose owner', trigger: 'change' }
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
  updated() {
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
          this.$confirm('Whether to update immediately?', 'tip', {
            confirmButtonText: 'confirm',
            cancelButtonText: 'cancel',
            type: 'success'
          }).then(() => {
            this.loading = true
            updateRealSequence({
              realSequenceID: this.accountId,
              owner: this.ruleForm.owner,
              totalLength: this.ruleForm.totalLength,
              dnaContents: this.ruleForm.dnaContents,
              description: this.ruleForm.description,
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: 'update success!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: 'update failed!'
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
      this.ruleForm.owner = accountId
    }
  }
}
</script>

<style scoped>
</style>
