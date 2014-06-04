# MTS timestamp

Go package for extracting timestamps from MTS video files

Install command-line tool with:

	go get github.com/homburg/mtstimestamp/extract_mts_timestamp

Usage:
	
	extract_mts_timestamp 00134.MTS

eg.

	extract_mts_timestamp 00134.MTS | xargs -I¤ rename "s/.*/¤.mts/" 00134.MTS

If a timestamp is not found, the filename (sans extension) is returned.

## LICENSE

MIT 2014 Thomas B Homburg
