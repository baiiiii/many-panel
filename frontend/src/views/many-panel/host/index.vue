<template>
    <div v-loading="loading">
        <LayoutContent title="主机">
            <template #toolbar>
                <el-row>
                    <el-col :span="16">
                        <el-button type="primary" @click="onOpenDialog('add')">添加主机</el-button>
                    </el-col>
                    <el-col :span="8">
                        <TableSetting @search="search()" />
                        <TableSearch @search="search()" v-model:searchName="searchName" />
                    </el-col>
                </el-row>
            </template>
            <template #main>
                <ComplexTable
                    :pagination-config="paginationConfig"
                    v-model:selects="selects"
                    :data="data"
                    @search="search"
                >
                    <el-table-column label="主机名" prop="name" min-width="60" />
                    <el-table-column label="主机地址" prop="host" min-width="60" />

                    <fu-table-operations width="200px" :buttons="buttons" :label="$t('commons.table.operate')" />
                </ComplexTable>
            </template>
        </LayoutContent>

        <OpDialog ref="opRef" @search="search" />
        <OperatorDialog @search="search" ref="dialogRef" />
    </div>
</template>

<script lang="ts" setup>
import OperatorDialog from '@/views/many-panel/host/operator/index.vue';
import { reactive, onMounted, ref } from 'vue';
import { MpHost } from '@/api/interface/mp-host';
import { searchMpHost, deleteMpHost, setDefaultHost } from '@/api/mp/mp-host';
import i18n from '@/lang';

const loading = ref();
const data = ref();
const selects = ref<any>([]);
const paginationConfig = reactive({
    cacheSizeKey: 'image-repo-page-size',
    currentPage: 1,
    pageSize: 10,
    total: 0,
});
const searchName = ref();

const opRef = ref();

const search = async () => {
    let params = {
        info: '',
        name: searchName.value,
        page: paginationConfig.currentPage,
        pageSize: paginationConfig.pageSize,
    };
    loading.value = true;
    await searchMpHost(params)
        .then((res) => {
            loading.value = false;
            data.value = res.data.items || [];
            paginationConfig.total = res.data.total;
        })
        .catch(() => {
            loading.value = false;
        });
};
const dialogRef = ref();
const onOpenDialog = async (title: string, rowData: Partial<MpHost.HostInfo> = {}) => {
    let params = {
        title,
        rowData: { ...rowData },
    };
    dialogRef.value!.acceptParams(params);
};

const onDelete = async (row: MpHost.HostInfo) => {
    let ids = [row.id];
    opRef.value.acceptParams({
        title: i18n.global.t('commons.button.delete'),
        names: [row.name],
        msg: i18n.global.t('commons.msg.operatorHelper', [i18n.global.t('commons.button.delete')]),
        api: deleteMpHost,
        params: { ids: ids },
    });
};

const onSetDefaultHost = async (row: MpHost.HostInfo) => {
    let id = row.id;
    opRef.value.acceptParams({
        title: '设为默认',
        names: [row.name],
        msg: '设为默认?',
        api: setDefaultHost,
        params: { id: id },
    });
};

const buttons = [
    {
        label: '设为默认',
        click: (row: MpHost.HostInfo) => {
            onSetDefaultHost(row);
        },
        disabled: (row: MpHost.HostInfo) => {
            return row.isDefault == '1';
        },
    },
    {
        label: i18n.global.t('commons.button.edit'),
        click: (row: MpHost.HostInfo) => {
            onOpenDialog('edit', row);
        },
    },
    {
        label: i18n.global.t('commons.button.delete'),
        click: (row: MpHost.HostInfo) => {
            onDelete(row);
        },
        disabled: (row: MpHost.HostInfo) => {
            return row.isDefault === '1';
        },
    },
];

onMounted(() => {
    search();
});
</script>
