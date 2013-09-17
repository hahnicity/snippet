package config

import "github.com/hahnicity/go-stringit"

var (
    DescPrefix           string = "DESCRIPTION: "
    DescQuery            string = "Please enter a brief description for this chunk of code:"
    DescSuffix           string = "\n"
    EndBlockSuffix       string = "\n\n"
    ImportantAffirmative string = "y"
    ImportantNegative    string = "N"
    ImportantQuery       string = stringit.Format(
        "Is this chunk of code important? ({}/{}): ",
        ImportantAffirmative, 
        ImportantNegative,
    )
    ImportantRetryQuery  string = stringit.Format(
        "Your response was not understood. Please enter {} OR {}: ",
        ImportantAffirmative,
        ImportantNegative,
    )
)
