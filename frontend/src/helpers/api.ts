export class API {
    /**
     * Request executes a http request and returns the response
     * @param url http destination url
     * @param token bearer token
     * @param method http method (POST, PUT etc...)
     * @param body body of request
     * @param headers additional headers of request
     */
    public static async Req(
        method: string,
        url: string,
        {body = null, headers = new Headers()}: { body?: BodyInit | null, headers?: Headers } = {}
    ): Promise<{ ok: boolean, response: any, status: number }> {
        const response = await fetch(url, {headers, method, body})
        const contentType = response.headers.get("content-type")
        return {
            response: contentType && contentType.indexOf("application/json") !== -1 ? await response.json() : await response.text(),
            ok: response.ok,
            status: response.status
        }
    }
}
