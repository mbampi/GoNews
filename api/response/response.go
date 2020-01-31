package response


// FIXME: not dealing errors correctly
func respondWith(w http.ResponseWriter, status int, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(data)
		//w.Write(js)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error: %v", err.Error())
	}
}