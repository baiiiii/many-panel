<template>
    <div v-loading="loading">
        <LayoutContent :divider="true">
            <template #main>
                <el-form :model="form" label-position="left" label-width="150px">
                    <el-row>
                        <el-col :span="1"><br /></el-col>
                        <el-col :xs="24" :sm="20" :md="15" :lg="12" :xl="12">
                            <el-form-item :label="$t('setting.user')" prop="userName">
                                <el-input disabled v-model="form.userName">
                                    <template #append>
                                        <el-button @click="onChangeUserName()" icon="Setting">
                                            {{ $t('commons.button.set') }}
                                        </el-button>
                                    </template>
                                </el-input>
                            </el-form-item>

                            <el-form-item :label="$t('setting.passwd')" prop="password">
                                <el-input type="password" disabled v-model="form.password">
                                    <template #append>
                                        <el-button icon="Setting" @click="onChangePassword">
                                            {{ $t('commons.button.set') }}
                                        </el-button>
                                    </template>
                                </el-input>
                            </el-form-item>

                            <el-form-item :label="$t('setting.sessionTimeout')" prop="sessionTimeout">
                                <el-input disabled v-model.number="form.sessionTimeout">
                                    <template #append>
                                        <el-button @click="onChangeTimeout" icon="Setting">
                                            {{ $t('commons.button.set') }}
                                        </el-button>
                                    </template>
                                </el-input>
                                <span class="input-help">
                                    {{ $t('setting.sessionTimeoutHelper', [form.sessionTimeout]) }}
                                </span>
                            </el-form-item>

                            <el-form-item :label="$t('setting.mfa')">
                                <el-switch
                                    @change="handleMFA"
                                    v-model="form.mfaStatus"
                                    active-value="enable"
                                    inactive-value="disable"
                                />
                                <span class="input-help">
                                    {{ $t('setting.mfaHelper') }}
                                </span>
                            </el-form-item>
                        </el-col>
                    </el-row>
                </el-form>
            </template>
        </LayoutContent>

        <Password ref="passwordRef" />
        <UserName ref="userNameRef" />
        <Timeout ref="timeoutRef" @search="search()" />
        <MfaSetting ref="mfaRef" @search="search" />
    </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue';
import { ElForm } from 'element-plus';
import { getSettingInfo, getSystemAvailable, updateSetting } from '@/api/mp/setting';
import { useI18n } from 'vue-i18n';
import Password from '@/views/many-panel/setting/password/index.vue';
import UserName from '@/views/many-panel/setting/username/index.vue';
import Timeout from '@/views/many-panel/setting/timeout/index.vue';
import MfaSetting from '@/views/many-panel/setting/mfa/index.vue';
import { MsgSuccess } from '@/utils/message';

const loading = ref(false);
const i18n = useI18n();

const form = reactive({
    userName: '',
    password: '',
    sessionTimeout: 0,
    complexityVerification: '',
    mfaStatus: 'disable',
    mfaInterval: 30,
});

const userNameRef = ref();
const passwordRef = ref();
const timeoutRef = ref();
const mfaRef = ref();

const search = async () => {
    const res = await getSettingInfo();
    form.userName = res.data.userName;
    form.password = '******';
    form.sessionTimeout = Number(res.data.sessionTimeout);
    form.complexityVerification = res.data.complexityVerification;
};

const onChangePassword = () => {
    passwordRef.value.acceptParams({ complexityVerification: form.complexityVerification });
};
const onChangeUserName = () => {
    userNameRef.value.acceptParams({ userName: form.userName });
};
const onChangeTimeout = () => {
    timeoutRef.value.acceptParams({ sessionTimeout: form.sessionTimeout });
};
const handleMFA = async () => {
    if (form.mfaStatus === 'enable') {
        mfaRef.value.acceptParams({ interval: form.mfaInterval });
        return;
    }
    loading.value = true;
    await updateSetting({ key: 'MFAStatus', value: 'disable' })
        .then(() => {
            loading.value = false;
            search();
            MsgSuccess(i18n.t('commons.msg.operationSuccess'));
        })
        .catch(() => {
            loading.value = false;
        });
};

onMounted(() => {
    search();
    getSystemAvailable();
});
</script>
