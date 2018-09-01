var resp = http.get(url, {"Accept": "application/json"});
if (!resp.Body) {
    throw error(resp)
}
system.call('testing1', {"hello": "Called internally"});
