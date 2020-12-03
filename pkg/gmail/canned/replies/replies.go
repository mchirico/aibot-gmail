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

Current Hourly Rate: $109/hr 

What's the hourly rate for this position? If it is
below $109/hr, I'll have to pass on it.

Please include your LinkedIn account. I'll attach a
more detailed work history, via LinkedIn messaging, if
it makes sense to move-to-the-next step.

I'll wait for your confirmation on these points. 

Please don't forget to include your LinkedIn account,
as we've never met, and I'd like to confirm who
you are.

It is not my intention to come off as rude with this 
direct approach.  I'm doing this out of respect for
your time and my time. Reaching an agreement, signing 
a contract, and coordinating work to solve your particular
requirement(s), is the only purpose of these communications.

Security and verification is important to me and the people
I work with.  Please understand this, and give me the resources
to verify who you are, as I only work with trusted clients.



Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only)
https://www.linkedin.com/in/mikechirico

* Please include mobile phone number
* Please include LinkedIn account (with photo)
* NO PHONE CALLS - Video Calls Only 
  `, `

Is this position 100% remote, and can this position work
on a corp-to-corp contract? If truly remote, please 
confirm that.

Do you have a mobile number, where I can reach you? Sorry,
but I'm not convinced you're a real person.

Current Hourly Rate: $109/hr 

What's the hourly rate for this position?


Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only)
`, `

For future reference: I'm only open to a corp-to-corp contract, with my
company CWXSTAT INC.  

Remote contract work ONLY.

Reaching an agreement, signing a contract, and coordinating work through
CWXSTAT is the purpose of these communications.   

For your reference, average contract hourly rate is between $115/hr to 
$130/hr, payable in 15 days.

Video calls are great and are preferred over just voice (Slack, Zoom, 
Skype etc.) so that diagrams can be shared, and technical ideas can be 
discussed. Are you okay with that? Video calls will be recorded so that
we can refer to points made later.

Please have direct knowledge of the position/work and requirements. If you 
are not the direct hiring manager, can you or the party you represent, have 
have that person present on the video call?  

So this is a little terse; but, under no circumstances may any information 
about CWXSTAT, or myself be represented, in any form, without prior written 
approval. Receiving a copy of a resume is NOT CONSENT to represent. Sorry, I
need to make that clear -- recieving a copy of a resume is NOT CONSENT to
represent. 


Regards,

Mike Chirico
mc@cwxstat.com
(215) 326-9389 (text only, for verification)
https://www.linkedin.com/in/mikechirico

1) Please be sure to include your LinkedIn account
2) Please include mobile number (text for security code)

`, `--000000000000eaa62105aeea3888
Content-Type: text/plain; charset="UTF-8"

For future reference: I'm only open to a corp-to-corp contract, with my
company CWXSTAT INC.

Remote contract work ONLY.

Hourly rate between 150/hr to 170/hr.

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
<p>Hourly rate between 150/hr to 170/hr.</p>
<p>Please be sure to include your mobile phone number, where you can be reached by text. A mobile phone is necessary for receiving emails in confidential mode.</p>
<p>Regards,</p>
<p>Mike Chirico<br><a href="mailto:mc@cwxstat.com">mc@cwxstat.com</a><br>(215) 326-9389 (text only)</p></div>

--000000000000eaa62105aeea3888--`,
	}
	return _list
}
