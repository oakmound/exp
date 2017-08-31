# Shiny

This is a fork of golang.org/exp/shiny. 
The goal of this fork is to add additional window management functionality, 
and otherwise actively maintain the project where it is not being maintained in
it's current form. 

## Long Term Plans 

Just window/screen and input management. The elements of shiny which relate to UI composition (flex boxes)
will be stripped off and put somewhere else, because we don't think its necessary for the window 
management unit to be aware of the structure of the UI being put into the window. 

If there are objections to this, feel free to raise them with use cases. 