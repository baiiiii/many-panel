<template>
    <div class="logo">
        <el-select v-model="value" placeholder="Select" size="large" class="host-select" @change="openHost">
            <el-option v-for="item in data" :key="item.id" :label="item.name" :value="item.id">
                <template #default>
                    <div class="flex items-center justify-between">
                        <div class="flex items-center gap-5">
                            <span>{{ item.name }}</span>
                        </div>
                        <div class="flex items-center gap-5">
                            <el-button @click.stop="openHostNewTab(item.id)">
                                <el-icon>
                                    <svg
                                        preserveAspectRatio="xMidYMid meet"
                                        viewBox="0 0 24 24"
                                        width="1.2em"
                                        height="1.2em"
                                        class="link-icon"
                                    >
                                        <path
                                            fill="currentColor"
                                            d="M10 6v2H5v11h11v-5h2v6a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h6zm11-3v8h-2V6.413l-7.793 7.794l-1.414-1.414L17.585 5H13V3h8z"
                                        ></path>
                                    </svg>
                                </el-icon>
                            </el-button>
                        </div>
                    </div>
                </template>
            </el-option>
        </el-select>
    </div>
</template>

<script setup lang="ts">
// import router from '@/routers';
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { LoginMpHost, searchMpHost } from '@/api/mp/mp-host';

defineProps<{ isCollapse: boolean }>();
const router = useRouter();
const value = ref();
const data = ref();
const search = async () => {
    let params = {
        info: '',
        page: 1,
        pageSize: 20,
    };
    await searchMpHost(params).then((res) => {
        data.value = res.data.items || [];
        if (sessionStorage.getItem('host-id')) {
            value.value = Number.parseInt(sessionStorage.getItem('host-id'));
        }
    });
};

const openHost = async (val: string) => {
    await LoginMpHost({ id: val }).then((res) => {
        if (res.code === 200) {
            history.go(0);
            sessionStorage.setItem('host-id', val);
        }
    });
};

const openHostNewTab = async (val: string) => {
    let page = router.resolve({
        path: '/',
        query: { id: val },
    });
    await LoginMpHost({ id: val }).then((res) => {
        if (res.code === 200) {
            window.open(page.href, '_blank');
        }
    });
};

onMounted(() => {
    search();
});
/*const goHome = () => {
    router.push({ name: 'home' });
};*/
</script>

<style scoped lang="scss">
.logo {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 55px;
    img {
        object-fit: contain;
        width: 95%;
        height: 45px;
    }
}
</style>

<style lang="scss">
.host-select {
    width: 180px;
    height: 53px;
    .el-select__wrapper {
        min-height: 53px !important;
        background-color: transparent !important;
    }
}
</style>
