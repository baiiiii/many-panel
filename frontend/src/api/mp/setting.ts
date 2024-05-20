import http from '@/api/mp';
import { Setting } from '../interface/setting';

export const getSettingInfo = () => {
    return http.post<Setting.SettingInfo>(`/settings/search`);
};
export const getSystemAvailable = () => {
    return http.get(`/settings/search/available`);
};

export const updateSetting = (param: Setting.SettingUpdate) => {
    return http.post(`/settings/update`, param);
};

export const updatePassword = (param: Setting.PasswordUpdate) => {
    return http.post(`/settings/password/update`, param);
};

export const loadMFA = (param: Setting.MFARequest) => {
    return http.post<Setting.MFAInfo>(`/settings/mfa`, param);
};

export const bindMFA = (param: Setting.MFABind) => {
    return http.post(`/settings/mfa/bind`, param);
};
