<template>
  <div class="container">
    <el-alert
      type="success"
    >
      <p>Account ID: {{ accountId }}</p>
      <p>username: {{ userName }}</p>
      <p>balance: ￥{{ balance }} yen</p>
      <p>When an assignment  or authorizing operation is initiated, the endorsement is true</p>
      <p>When the endorsement is false, the assignment or authorizing operation can only be initiated</p>
    </el-alert>
    <div v-if="realSequenceList.length==0" style="text-align: center;">
      <el-alert
        title="can not find the information"
        type="warning"
      />
    </div>
    <el-row v-loading="loading" :gutter="20">
      <el-col v-for="(val,index) in realSequenceList" :key="index" :span="6" :offset="1">
        <el-card class="realSequence-card">
          <div slot="header" class="clearfix">
            Endorse status:
            <span style="color: rgb(255, 0, 0);">{{ val.endorsement }}</span>
          </div>

          <div class="item">
            <el-tag>Account ID: </el-tag>
            <span>{{ val.realSequenceId }}</span>
          </div>
          <div class="item">
            <el-tag type="success">patient ID: </el-tag>
            <span>{{ val.patient }}</span>
          </div>
          <div class="item">
            <el-tag type="warning">total length: </el-tag>
            <span>{{ val.totalLength }}</span>
          </div>
          <div class="item">
            <el-tag type="danger">dna contents: </el-tag>
            <span>{{ val.dnaContents }}</span>
          </div>

          <div v-if="!val.endorsement&&roles[0] !== 'admin'">
            <el-button type="text" @click="openDialog(val)">authorizing</el-button>
            <el-divider direction="vertical" />
            <el-button type="text" @click="openAppointingDialog(val)">appointing</el-button>
          </div>
          <el-rate v-if="roles[0] === 'admin'" />
        </el-card>
      </el-col>
    </el-row>
    <el-dialog v-loading="loadingDialog" :visible.sync="dialogCreateAuthorizing" :close-on-click-modal="false" @close="resetForm('realForm')">
      <el-form ref="realForm" :model="realForm" :rules="rules" label-width="100px">
        <el-form-item label="price (yen)" prop="price">
          <el-input-number v-model="realForm.price" :precision="2" :step="10000" :min="0" />
        </el-form-item>
        <el-form-item label="Expiry date (days)" prop="authorizePeriod">
          <el-input-number v-model="realForm.authorizePeriod" :min="1" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="createAuthorizing('realForm')">authorizing immediately</el-button>
        <el-button @click="dialogCreateAuthorizing = false">cancel</el-button>
      </div>
    </el-dialog>
    <el-dialog v-loading="loadingDialog" :visible.sync="dialogCreateAppointing" :close-on-click-modal="false" @close="resetForm('AppointingForm')">
      <el-form ref="AppointingForm" :model="AppointingForm" :rules="rulesAppointing" label-width="100px">
        <el-form-item label="patient" prop="patient">
          <el-select v-model="AppointingForm.patient" placeholder="please choose patient" @change="selectGet">
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
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="createAppointing('AppointingForm')">appointing immediately</el-button>
        <el-button @click="dialogCreateAppointing = false">cancel</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { queryRealSequenceList } from '@/api/realSequence'
import { createAuthorizing } from '@/api/authorizing'
import { createAppointing } from '@/api/appointing'

export default {
  name: 'RealSequence',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('must be bigger than0'))
      } else {
        callback()
      }
    }
    return {
      loading: true,
      loadingDialog: false,
      realSequenceList: [],
      dialogCreateAuthorizing: false,
      dialogCreateAppointing: false,
      realForm: {
        price: 0,
        authorizePeriod: 0
      },
      rules: {
        price: [
          { validator: checkArea, trigger: 'blur' }
        ],
        authorizePeriod: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      AppointingForm: {
        patient: ''
      },
      rulesAppointing: {
        patient: [
          { required: true, message: 'please choose patient', trigger: 'change' }
        ]
      },
      accountList: [],
      valItem: {}
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ])
  },
  created() {
    if (this.roles[0] === 'admin') {
      queryRealSequenceList().then(response => {
        if (response !== null) {
          this.realSequenceList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryRealSequenceList({ patient: this.accountId }).then(response => {
        if (response !== null) {
          this.realSequenceList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
    openDialog(item) {
      this.dialogCreateAuthorizing = true
      this.valItem = item
    },
    openAppointingDialog(item) {
      this.dialogCreateAppointing = true
      this.valItem = item
      queryAccountList().then(response => {
        if (response !== null) {
          // filter admin and current user
          this.accountList = response.filter(item =>
            item.userName !== 'admin' && item.accountId !== this.accountId
          )
        }
      })
    },
    createAuthorizing(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('Whether authorizing immediately?', 'tip', {
            confirmButtonText: 'confirm',
            cancelButtonText: 'cancel',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createAuthorizing({
              objectOfAuthorize: this.valItem.realSequenceId,
              patient: this.valItem.patient,
              price: this.realForm.price,
              authorizePeriod: this.realForm.authorizePeriod
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateAuthorizing = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: 'authorizing success!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: 'authorizing failed!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateAuthorizing = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateAuthorizing = false
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
    createAppointing(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('Whether to donate immediately?', 'tip', {
            confirmButtonText: 'confirm',
            cancelButtonText: 'cancel',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createAppointing({
              objectOfAppointing: this.valItem.realSequenceId,
              patient: this.valItem.patient,
              hospital: this.AppointingForm.patient
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateAppointing = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: 'appointing success!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: 'appointing failed!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateAppointing = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateAppointing = false
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
      this.AppointingForm.patient = accountId
    }
  }
}

</script>

<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .realSequence-card {
    width: 280px;
    height: 340px;
    margin: 18px;
  }
</style>
