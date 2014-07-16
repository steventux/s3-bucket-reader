go-cloud-gallery
================

Retrieve cloud file storage URLs with Go.

Very much a proof of concept, the aim of this project is to provide cloud storage URLs from a number of services.
For now this prototype is hardwired to AWS/S3 USEast storage buckets but the intention is to provide multiple service adapters with no regional hardwiring.
Ultimately this would grow into a RESTful API providing JSON responses for use with common frontend image and video gallery libraries. 
