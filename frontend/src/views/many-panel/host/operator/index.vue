<template>
    <el-drawer v-model="drawerVisible" :destroy-on-close="true" :close-on-click-modal="false" size="40%">
        <template #header>
            <DrawerHeader :header="title" :resource="dialogData.rowData?.name" :back="handleClose" />
        </template>
        <el-form
            ref="formRef"
            label-position="top"
            v-loading="loading"
            :model="dialogData.rowData"
            :rules="rules"
            label-width="120px"
        >
            <el-row type="flex" justify="center">
                <el-col :span="22">
                    <el-form-item label="主机名" prop="name">
                        <el-input clearable v-model.trim="dialogData.rowData!.name"></el-input>
                    </el-form-item>
                    <el-form-item label="主机地址" prop="host">
                        <el-input
                            clearable
                            v-model.trim="dialogData.rowData!.host"
                            :placeholder="'http://127.0.0.1:8080'"
                        ></el-input>
                    </el-form-item>
                    <el-form-item label="用户名" prop="user">
                        <el-input clearable v-model.trim="dialogData.rowData!.user"></el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="pwd">
                        <el-input
                            clearable
                            type="password"
                            show-password
                            v-model.trim="dialogData.rowData!.pwd"
                        ></el-input>
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button :disabled="loading" @click="drawerVisible = false">
                    {{ $t('commons.button.cancel') }}
                </el-button>
                <el-button :disabled="loading" type="primary" @click="onSubmit(formRef)">
                    {{ $t('commons.button.confirm') }}
                </el-button>
            </span>
        </template>
    </el-drawer>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue';
import { Rules } from '@/global/form-rules';
import i18n from '@/lang';
import { ElForm } from 'element-plus';
import { MpHost } from '@/api/interface/mp-host';
import DrawerHeader from '@/components/drawer-header/index.vue';
import { createMpHost, updateMpHost } from '@/api/mp/mp-host';
import { MsgSuccess } from '@/utils/message';

const loading = ref(false);

interface DialogProps {
    title: string;
    rowData?: MpHost.HostInfo;
    getTableList?: () => Promise<any>;
}
const title = ref<string>('');
const drawerVisible = ref(false);
const dialogData = ref<DialogProps>({
    title: '',
});
const acceptParams = (params: DialogProps): void => {
    dialogData.value = params;
    title.value = i18n.global.t('commons.button.' + dialogData.value.title);
    drawerVisible.value = true;
};
const emit = defineEmits<{ (e: 'search'): void }>();

const handleClose = () => {
    drawerVisible.value = false;
};
const rules = reactive({
    name: [Rules.requiredInput],
    host: [Rules.requiredInput],
    user: [Rules.requiredInput],
    pwd: [Rules.requiredInput],
});

type FormInstance = InstanceType<typeof ElForm>;
const formRef = ref<FormInstance>();

const onSubmit = async (formEl: FormInstance | undefined) => {
    if (!formEl) return;
    formEl.validate(async (valid) => {
        if (!valid) return;
        loading.value = true;
        if (dialogData.value.title === 'add') {
            await createMpHost(dialogData.value.rowData!)
                .then(() => {
                    loading.value = false;
                    MsgSuccess(i18n.global.t('commons.msg.operationSuccess'));
                    emit('search');
                    drawerVisible.value = false;
                })
                .catch(() => {
                    loading.value = false;
                });
            return;
        }
        await updateMpHost(dialogData.value.rowData!)
            .then(() => {
                loading.value = false;
                MsgSuccess(i18n.global.t('commons.msg.operationSuccess'));
                emit('search');
                drawerVisible.value = false;
            })
            .catch(() => {
                loading.value = false;
            });
    });
};

defineExpose({
    acceptParams,
});
</script>
