### Vagrant setup

1. Install [Vagrant](http://www.vagrantup.com/downloads) for your platform.
1. Open a terminal window, and cd /path/to/this/repository.

  ```bash
  # open a terminal window and type:
  $ cd golang-lab-slack
  ```

1. From within the `golang-lab-slack` directory, type `vagrant up`. This will create a virtual machine with Ubuntu linux and Go 1.4 installed for the Vagrant user.
1. Type `vagrant ssh` to ssh into the virtual machine.  

Within the virtual machine, `golang-lab-slack` on the host computer
will be mapped to /opt/golang-lab-slack in the guest.  So any changes
you make in the normal filesystem should be reflected in the VM!

Need help? Arrive a little early and we can help you get up-and-running, or join
the Boston Golang Slack group #lab-help channel group by signing up [here](http://bostongolang-slack-invite.herokuapp.com/).

You can also email me directly: [jandre+bostongolang@gmail.com](mailto:jandre+bostongolang@gmail.com).


