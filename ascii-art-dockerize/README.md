Description:
This following application's back end is written in Go and the front end is written in html. 
The application takes input text and a font from the user and outputs an ascii art representation of the text  

In this web application you can use different fonts of ASCII-ART to print what you wish.

https://git.learn.01founders.co/Jonnytan555/ascii-art-web.git

Author:
Jonathan Edwards

Usage: 
1. Open up a browser and search http://localhost:8080/
2. Enter the text that you want to be converted into ASCII art into the top box.
3. Chhose which banner you would like to use.
4. Click on submit button.
5. Your text will be displayed in the box below.

Implementation Details: 
We used bufio to append each line to an array.
We used map function to store lines for every character.