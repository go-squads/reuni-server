package services

func createServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return createService(serviceStore)
}

func deleteServiceProcess(servicedata servicev) error {
	serviceStore := service{}
	serviceStore.Name = servicedata.Name
	return deleteService(serviceStore)
}
