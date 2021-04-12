#!/bin/sh -l
# Add the path to the profile running this sh
PATH=$PATH:/bin:/usr/bin:/go/bin
export PATH
# Check out the environment variables (for debugging)
echo "Environment Variables"
printenv
echo "Action Args Passed"
sh -c "echo $*"
# Autoversion
echo "Checking Autoversion"
PREVIOUS_VERSION=$(autoversion version previous .)
NEXT_VERSION=$(autoversion version next .)
echo "Previous Version: $PREVIOUS_VERSION"
echo "Next Version: $NEXT_VERSION"