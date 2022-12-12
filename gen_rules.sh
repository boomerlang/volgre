#!/bin/sh


for (( i=1; i<=$1; i++ )); do
echo "rule CheckCardNumber$i \"check card number structure $i\"  salience 10 {
    when
		CreditCard.ValidateCardNumber()
    then
		ValidationResults.CardNumber = true;
		Retract(\"CheckCardNumber$i\");
}

rule Check1$i \"test chaining $i\"  salience 10 {
	when
		ValidationResults.CardNumber == true
	then
		Log(\"Check1 = TRUE\");
		Retract(\"Check1$i\");
}

rule CheckHolderName$i \"check holder name $i\"  salience 10 {
    when
        CreditCard.ValidateHolderName() == true
    then
        ValidationResults.OwnerName = true;
        Retract(\"CheckHolderName$i\");
}

rule CheckExpMonth$i \"check expiration month $i\"  salience 10 {
    when
        CreditCard.ValidateExpMonth() == true
    then
        ValidationResults.ExpireMonth = true;
        Retract(\"CheckExpMonth$i\");
}

rule CheckExpYear$i \"check expiration year $i\"  salience 10 {
	when
		CreditCard.ValidateExpYear() == true
	then
	    ValidationResults.ExpireYear = true;
	    Retract(\"CheckExpYear$i\");
}

rule CheckSecurityCode$i \"check security code $i\" salience 10 {
	when
		CreditCard.ValidateSecurityCode() == true
	then
	    ValidationResults.SecurityCode = true;
	    Retract(\"CheckSecurityCode$i\");
}
"

done
