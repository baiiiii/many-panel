<template>
    <div>
        <RouterButton :buttons="buttons" />
        <LayoutContent>
            <router-view></router-view>
        </LayoutContent>
    </div>
</template>

<script lang="ts" setup>
const buttons = [
    {
        label: '主机',
        path: '/mp/host',
    },
    {
        label: '设置',
        path: '/mp/setting',
    },
];
</script>
