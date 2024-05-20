export namespace MpHost {
    export interface HostInfo {
        id: number;
        createdAt: Date;
        name: string;
        host: string;
        user: string;
        pwd: string;
        isDefault: string;
    }

    export interface HostCreate {
        name: string;
        host: string;
        user: string;
        pwd: string;
    }
    export interface HostUpdate {
        id: number;
        name: string;
        host: string;
        user: string;
        pwd: string;
    }
    export interface HostDelete {
        ids: Array<number>;
    }
}
