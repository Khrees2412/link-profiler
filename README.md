  This is a simple CLI tool built to solve the [Systems Engineering Assignment](https://github.com/cloudflare-hiring/cloudflare-2020-systems-engineering-assignment) from Cloudflare.

# Installation
- Ensure you have the Go [compiler](https://go.dev) installed.
- Run in terminal
`go install github.com/khrees2412/linkcli`
# Usage
#### To get the profile of a URL simply run 
 - `./linkcli get --url=<the URL or Site Link> --profile=<the number of times to evaluate the link`
& e.g  `./linkcli get --url=https://google.com --profile=6`
#### For multiple URLs or links, simply separate them with a comma
e.g `./linkcli get --url=https://google.com,https://github.com,https://slack.com --profile=3`

#### An output explainer is shown below:
All values for fastest, slowest, mean and median are in seconds.
```json

{
  "fastest": "<This is the record in seconds of the fastest request amount the number of requests you profiled>",
  "slowest": "<This is the record in seconds of the slowest request amount the number of requests you profiled>",
  "mean": "<This is the mean of all the requests you profiled>",
  "median": "<This is the median of all the requests you profiled>",
  "percentageSuccessful": "<This is the percentage of requests that succeeded>",
  "errorCodes": "[<This is a list containing all error codes returned>]", 
  "bytesOfSmallest": "<This is the size of the smallest response>",
  "bytesOfLargest": "<This is the size of the largest response>"
}
```

# Contributing

Simply create an issue, make a PR and link the issue. I'll review and merge. Thank you!

**Don't forget to leave a star**