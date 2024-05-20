import http from '@/api/mp';
import { TimeoutEnum } from '@/enums/http-enum';
import { ResPage, SearchWithPage } from '@/api/interface';
import { MpHost } from '@/api/interface/mp-host';

export const searchMpHost = (params: SearchWithPage) => {
    return http.post<ResPage<MpHost.HostInfo>>(`/host/search`, params);
};
export const createMpHost = (params: MpHost.HostCreate) => {
    return http.post(`/host/create`, params, TimeoutEnum.T_40S);
};
export const updateMpHost = (params: MpHost.HostUpdate) => {
    return http.post(`/host/update`, params, TimeoutEnum.T_40S);
};
export const deleteMpHost = (params: MpHost.HostDelete) => {
    return http.post(`/host/delete`, params, TimeoutEnum.T_40S);
};
export const setDefaultHost = (params: { id: string }) => {
    return http.get(`/host/default/${params.id}`);
};
export const LoginMpHost = (params: { id: string }) => {
    return http.get(`/host/login/${params.id}`);
};
