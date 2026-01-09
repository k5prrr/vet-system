const ajax = {
    versionDate: '2026-01-09',

    handleCallback(callback, data, arg) {
        if (typeof callback !== 'function')
            callback = console.log

        if (data === undefined)
            callback()
        else if (arg === undefined)
            callback(data)
        else
            callback(data, arg)
    },

    fetchJson(url, options = {}, callback = console.log, arg) {
        fetch(url, options)
            .then(response => {
                const contentType = response.headers.get('content-type')
                if (contentType && contentType.includes('application/json')) {
                    return response.json().catch(() => ({
                        error: "Invalid JSON",
                        raw: "(failed to parse JSON)"
                    }))
                } else {
                    return response.text().then(text => ({
                        error: "Non-JSON response",
                        raw: text
                    }))
                }
            })
            .catch(networkError => ({
                error: "Network error",
                message: networkError.message || "Failed to fetch"
            }))
            .then(result => {
                ajax.handleCallback(callback, result, arg)
            })
    },

    json(url, data = {}, callback = console.log, arg) {
        if (!url) throw new Error('URL is required.')

        let method = 'POST'
        if (data.ajaxMethod) {
            method = data.ajaxMethod
            delete data.ajaxMethod
        }

        if (method === 'GET')
            return ajax.getJson(url, callback, arg)

        ajax.fetchJson(url, {
            method,
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        }, callback, arg)
    },

    getJson(url, callback = console.log, arg) {
        if (!url) throw new Error('URL is required.')
        ajax.fetchJson(url, {}, callback, arg)
    },

    text(url, callback = console.log, arg) {
        if (!url) throw new Error('URL is required.')
        fetch(url)
            .then(response => response.text())
            .then(text => ajax.handleCallback(callback, text, arg))
            .catch(error => {
                const errObj = { error: "Network error", message: error.message }
                ajax.handleCallback(callback, errObj, arg)
            });
    },

    sendForm(url, form, callback = console.log, arg) {
        if (!url) throw new Error('URL is required.')
        const formData = new FormData(form)
        ajax.fetchJson(url, {
            method: 'POST',
            body: formData
        }, callback, arg)
    }
}