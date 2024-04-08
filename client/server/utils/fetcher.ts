export const fetcher = async (url: string, method: string, body: any, token: string) => {
    try {
        const response = await fetch(url, {
            method: method,
            body,
            headers: {
                Authorization: `Bearer ${token}`,
            },
        });
        return response.json();
    } catch (e) {
        return {
            status: 500,
            body: 'Internal server error',
        };
    }
}