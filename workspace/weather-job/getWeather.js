var resp = http.get(url);
if (!resp.Body) {
    throw error(resp)
}
system.call('testing1', {"hello": "Called internally"});
