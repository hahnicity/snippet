package config


var (
    DescPrefix           string = "DESCRIPTION: "
    DescQuery            string = "Please enter a brief description for this chunk of code:"
    DescSuffix           string = "\n"
    EndBlockSuffix       string = "\n\n"
    ImportantAffirmative string = "y"
    ImportantNegative    string = "N"
    ImportantQuery       string = "Is this chunk of code important? ("+
        ImportantAffirmative+"/"+ImportantNegative+"): "
    ImportantRetryQuery  string = "Your response was not understood."+
        " please enter "+ImportantAffirmative+" OR "+ImportantNegative+": "
)
