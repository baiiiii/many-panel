// new WebSocket( -> newWebSocket(

export const newWebSocket = (url) => {
    const separator = url.includes('?') ? '&' : '?';
    return new WebSocket(`${url}${separator}host-id=${sessionStorage.getItem('host-id')}`);
};
