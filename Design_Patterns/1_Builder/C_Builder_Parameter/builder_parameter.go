package builderparameter

import "strings"

/*
One question that might be asked is how do you get the uses of the API to actually use the builders,
as opposed to manipulate the objects directly.
One approach is to simply hide the objects that you want your users not to touch.
*/

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("from email should contain @")
	}

	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("destination email should contain @")
	}

	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	// ...
}

/*
So now we come to the important part: How do we actually get people to send the email?
Because obviously somewhere behind the scenes, we want to have a function called Send Email Impl,
for example, which actually takes the email and it does whatever you need to do with it,
but we don't want our clients to actually work with the email object.
We want to work with a builder.
*/
// Declaring a build parameter, that is basically a function which applies to the builder
type build func(*EmailBuilder)

// We are going to use this build function in the publicly exposed function called send email
// SendEmail is the function people are meant to be using.
// It takes an argument action of type build.
// What happens is that whenever somebody calls  the function, they have to provide the body of a function,
// which takes an email builder as the first and only parameter and doesn't return any values
func SendEmail(action build) {
	// Initializing the builder
	builder := EmailBuilder{}
	// Builder pointer is passed as an argument for the action.
	action(&builder)
	// Then we would do the internal functioning of the sendMailImpl
	sendMailImpl(&builder.email)
}

func main() {
	// From the client's perspective, they would have to call the SendEmail function.
	//
	SendEmail(func(b *EmailBuilder) {
		b.
			From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
	/*
	 So, what happens is that when we call this function we create a builder, whichs is an email builder.
	 Then, we apply the action, which is the entire body of the SendEmail function.
	 So we invoke this entire function (action) on the builder pointer.
	 By the time we do the impl function, the builder has been initialized.
	 So when we call SendMailImpl, taking the email object, that object that the clients
	 are not supposed to see, we have everything initialized correctly with all the validations
	 and all the rest of it.
	*/
}
