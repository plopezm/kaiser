var resp = http.get(url);
if (!resp.Body) {
    throw error(resp)
}