package replies

type Replies struct {
	Id     int
	NextId int
}

func (r *Replies) Next() string {
	i := r.NextId
	r.NextId += 1
	if i >= 0 && i < len(g()) {
		r.Id = i
		return g()[r.Id]
	}
	r.Id = len(g()) - 1
	r.NextId = 0
	return g()[r.Id]
}

func (r *Replies) Get(i int) string {
	if i >= 0 && i < len(g()) {
		r.Id = i
		return g()[r.Id]
	}
	r.Id = len(g()) - 1
	return g()[r.Id]
}

func g() []string {
	_list := []string{
		`

Please confirm the position is 100% remote, and can 
work on a corp-to-corp contract. Please confirm this.

Current Hourly Rate: $105/hr 

What's the hourly rate for this position? If it is
below $105/hr, I'll have to pass on it.

Please include your LinkedIn account. I'll attach a
more detailed work history, via LinkedIn messaging, if
it makes sense to move-to-the-next step.

I'll wait for your confirmation on these points. 

Please don't forget to include your LinkedIn account,
as we've never met, and I'd like to confirm who
you are.



Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only)
https://www.linkedin.com/in/mikechirico

* Please include mobile phone number
* Please include LinkedIn account
`, `

Is this position 100% remote, and can this position work
on a corp-to-corp contract? If truly remote, please 
confirm that.

Do you have a mobile number, where I can reach you? Sorry,
but I'm not convinced you're a real person.

Current Hourly Rate: $105/hr 

What's the hourly rate for this position?


Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only)
`, `

For future reference: I'm only open to a corp-to-corp contract, with my
company CWXSTAT INC.

Remote contract work ONLY.

Hourly rate between $105/hr to $117/hr.

Please be sure to include your mobile phone number, where you can be
reached by text. A mobile phone is necessary for receiving emails in
confidential mode.  

Also, please include your LinkedIn account. This way, I have background
on you and your company, so that we can both move towards a possible
contract agreement, which is the only purpose of these emails.



Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only)
https://www.linkedin.com/in/mikechirico

1) Please be sure to include your LinkedIn account
2) Please include mobile number (text for security code)

`, `--000000000000eaa62105aeea3888
Content-Type: text/plain; charset="UTF-8"

For future reference: I'm only open to a corp-to-corp contract, with my
company CWXSTAT INC.

Remote contract work ONLY.

Hourly rate between 98/hr to 117/hr.

Please be sure to include your mobile phone number, where you can be
reached by text. A mobile phone is necessary for receiving emails in
confidential mode.

Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only)

--000000000000eaa62105aeea3888
Content-Type: text/html; charset="UTF-8"

<div dir="ltr"><p>For future reference: I&#39;m only open to a corp-to-corp contract, with my company CWXSTAT INC.  </p>
<p>Remote contract work ONLY.  </p>
<p>Hourly rate between 98/hr to 117/hr.</p>
<p>Please be sure to include your mobile phone number, where you can be reached by text. A mobile phone is necessary for receiving emails in confidential mode.</p>
<p>Regards,</p>
<p>Mike Chirico<br><a href="mailto:mc@cwxstat.com">mc@cwxstat.com</a><br>(215) 326-9389 (text only)</p></div>

--000000000000eaa62105aeea3888--`,
	}
	return _list
}
