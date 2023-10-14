package packages

import (
	"Nosviak2/core/clients/sessions"
	"Nosviak2/core/configs"
	"Nosviak2/core/sources/language/evaluator"
	"Nosviak2/core/sources/language/lexer"
	"Nosviak2/core/sources/tools"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func init() {

	RegisterPackage(evaluator.Package{
		Package: "colour",
		Functions: map[string]evaluator.Builtin{
			"hex" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the length given properly
				//this will ensure its done without any errors
				if len(args) != 1 { //returns the error properly without issues
					return make([]evaluator.Object, 0), errors.New("format: hex string")
				}
				//parses the input without issues
				//this will make sure its done without any errors
				vals, err := strconv.ParseUint(args[0].Literal(), 16, 32)
				if err != nil { //err handles properly without issues
					return make([]evaluator.Object, 0), err //error returns
				}
				R := uint8(vals >> 16) //Red
				G := uint8((vals >> 8) & 0xFF) //Green
				B := uint8(vals & 0xFF) //Blue
				//returns the values properly
				//this will ensure its done without any errors
				return evaluator.ArrayObject(evaluator.Object{Literal: strconv.Itoa(int(R))+","+strconv.Itoa(int(G))+","+strconv.Itoa(int(B)), Type: lexer.String}), nil
			},

			"rgb" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				//checks the length without issues happening on purpose
				if len(args) != 3 { //returns the error properly without issues
					return make([]evaluator.Object, 0), errors.New("format: r int, g int, b int")
				}
				//r sum for the statement
				r, err := strconv.Atoi(args[0].Literal()) //r int
				if err != nil || args[0].TokenType() != lexer.Int { //checks the values properly
					return make([]evaluator.Object, 0), errors.New("format: r int, g int, b int")
				}
				//g sum for the statement
				g, err := strconv.Atoi(args[1].Literal()) //g int
				if err != nil || args[1].TokenType() != lexer.Int { //checks the values properly
					return make([]evaluator.Object, 0), errors.New("format: r int, g int, b int")
				}
				//b sum for the statement
				b, err := strconv.Atoi(args[2].Literal()) //r int
				if err != nil || args[2].TokenType() != lexer.Int { //checks the values properly
					return make([]evaluator.Object, 0), errors.New("format: r int, g int, b int")
				}
				//returns the values properly
				//this will ensure its done without any errors
				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: strconv.Itoa(r)+","+strconv.Itoa(g)+","+strconv.Itoa(b)}), nil
			},

			"marshal" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {
				if len(args) < 3 { //checks the length correctly and properly
					return make([]evaluator.Object, 0), errors.New("missing arguments inside function call")
				}

				//gets the text properly
				//this will store the text items properly without issues
				Text := args[0].Literal() //stores the text properly
				Colours := make([]string, 0) //stores all the colours properly

				//ranges through the args properly
				//this will make sure its done correctly
				for _, arg := range args[1:] {
					//saves into the array of colours properly
					Colours = append(Colours, arg.Literal())
				}

				//performs the gradient properly
				//this will make sure its done properly without issues
				//output, err := tools.MakeGradient(Text, Colours, "*", make(map[string]int)).Perform([]string{Text}, "")
				//if err != nil { //basic error handling properly without issues
				//	return make([]evaluator.Object, 0), err
				//}

				output := Text

				return evaluator.ArrayObject(evaluator.Object{Type: lexer.String, Literal: output+"\x1b[0m"}), nil
			},


			//builtin gif client within the system properly
			//this will ensure its done without any errors happening
			"gif" : func(args []lexer.Token, s *sessions.Session, e *evaluator.Evaluator, wr io.Writer) ([]evaluator.Object, error) {

				//checks the amount of arguments provided
				//this will ensure its done without errors happening
				if len(args) < 4 { //ensures amount of arguments wanted
					return make([]evaluator.Object, 0), errors.New("missing pathway to gif target")
				}

				//tries to open the file properly
				//this will ensure its done without errors
				Gif, err := os.Open(filepath.Join(deployment.Assets, args[0].Literal()))
				if err != nil { //error handles properly and safely
					return make([]evaluator.Object,
						 0), err
				}

				//decodes the incoming frames from the gif properly
				//this will ensure its done without errors happening
				ProperImageLayers, err := gif.DecodeAll(Gif) //decodes all frames
				if err != nil { //error handles properly
					return make([]evaluator.Object, 0), err
				}

				//gets the images properly and safely properly
				//this will ensure its done without errors happening
				width, height := getGifDimensions(ProperImageLayers)
				overPaintImage := image.NewRGBA(image.Rect(0, 0, width, height))
				draw.Draw(overPaintImage, overPaintImage.Bounds(), ProperImageLayers.Image[0], image.Point{0,0}, draw.Src)


				//part of the system resize
				//this will ensure its done without errors
				DefaultX, DefaultY := s.Length, s.Height

				//converts into the int format properly
				//this will ensure its done without errors
				WantedX, err := strconv.Atoi(args[2].Literal())
				if err != nil { //error handles properly
					return make([]evaluator.Object, 0), err
				}

				//converts into the int format properly
				//this will ensure its done without errors
				WantedY, err := strconv.Atoi(args[1].Literal())
				if err != nil { //error handles properly
					return make([]evaluator.Object, 0), err
				}

				//converts into the int format properly
				//this will ensure its done without errors
				Loops, err := strconv.Atoi(args[3].Literal())
				if err != nil { //error handles properly
					return make([]evaluator.Object, 0), err
				}

				s.Write("\033c") //clears screen

				//ranges through the amount of times in loops
				//this will ensure its done without errors happening
				for proc := 0; proc < Loops; proc++ { //loops through properly

					var location int = 0
					//ranges through every frame properly and safely
					//this will ensure its done without errors happening
					for _, imageFrame := range ProperImageLayers.Image { //ranges
						draw.Draw(overPaintImage, overPaintImage.Bounds(), imageFrame, image.Point{0,0}, draw.Over)
						frame := image.NewRGBA(image.Rect(0, 0, width, height)) //creates the image properly
						draw.Draw(frame, frame.Bounds(), overPaintImage, image.Point{0,0}, draw.Over) //draws over properly
						
						//renders into the image scale properly and safely
						//this will ensure its done without errors happening on purpose
						converted, err := tools.NewScaledFromImage(frame, WantedY*2, WantedX, color.Black, 0, 0)
						if err != nil { //error handles properly
							log.Printf("[PACKAGES] [GIF]: %s\r\n", err.Error())
						}
	
						//writes the raw converted image
						//this will ensure its done without errors happening
						if err := s.Write("\x1b[?25l\x1b[0;0H\x1b[8;"+strconv.Itoa(WantedY)+";"+strconv.Itoa(WantedX)+"t\x1b[0;0H"+converted.Render()+"\x1b[0J"); err != nil {
							log.Printf("[PACKAGES] [GIF]: %s\r\n", err.Error())
						}
	
						//continues looping properly and safely
						//this will ensure its done without errors
						time.Sleep(time.Duration(ProperImageLayers.Delay[location]) * time.Millisecond); location++
					}
				}

				//sets the terminal size back to normal
				//this will ensure its done without errors
				if err := s.Write("\x1b[8;"+strconv.Itoa(DefaultY)+";"+strconv.Itoa(DefaultX)+"t\x1b[?25h\x1b[0;00M"); err != nil {
					return make([]evaluator.Object, 0), err //returns the information
				}

				return make([]evaluator.Object, 0), nil
			},
		},
	})
}


//gets the gifs largest dimensions within the frames
//this will ensure its done without errors happening on pupose
func getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX int //lowest recorded size X
	var lowestY int //lowest recorded size Y
	var highestX int //highest recorded size X
	var highestY int //highest recorded size Y

	//ranges through the frames properly
	//this will ensure its done without errors
	for _, img := range gif.Image { //ranges through properly
		if img.Rect.Min.X < lowestX { //checks safely
			lowestX = img.Rect.Min.X //gets the lowest
		}
		if img.Rect.Min.Y < lowestY { //checks safely
			lowestY = img.Rect.Min.Y //gets the lowest
		}
		if img.Rect.Max.X > highestX { //checks safely
			highestX = img.Rect.Max.X //gets the highest
		}
		if img.Rect.Max.Y > highestY { //checks safely
			highestY = img.Rect.Max.Y //gets the highest
		}
	}

	//returns the values properly without errors
	//this will ensure its done without errors happening
	return highestX - lowestX, highestY - lowestY
}
