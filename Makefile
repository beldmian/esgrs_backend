gen:
	microgen -main -file data_collection_service/service.go -out data_collection_service -package esgrs/data_collection_service
	microgen -main -file data_processing_service/service.go -out data_processing_service -package esgrs/data_processing_service
	microgen -main -file reporting_service/service.go -out reporting_service -package esgrs/reporting_service
	microgen -main -file visualization_service/service.go -out visualization_service -package esgrs/visualization_service

run:
	docker-compose up