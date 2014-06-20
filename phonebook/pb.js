var path = require('path'),
	fs = require('fs')

__dirname = process.cwd()

var taffy = require('taffy')

var Parser = function (){
	this.commands = {},
	this.defaultCommand = null
}

Parser.prototype.addCommand = function (command, callback) {

	this.commands[command] = callback
}

Parser.prototype.setDefaultCommand = function (command) {

	if (command in this.commands){

		this.defaultCommand = command
	}
	else {

		throw Error("Default command is not in command list")
	}
	
}

Parser.prototype.handleCommand =  function (command, args){

	if (command in this.commands){

		this.commands[command](args)
	}
	else if (this.defaultCommand){

		this.commands[this.defaultCommand](args)
	}
}
Parser.prototype.start = function (args) {

	if (args.length > 0){

		this.handleCommand(args[0], args)
	}
	else {

		//TODO:Better error handling
		console.log("No arguments")
	}
}

var Phonebook = {

	createFile: function (name, callback){

		var file = path.join(__dirname, name)
		var initialData = JSON.stringify([])

		fs.writeFile(file, initialData, function (err){

			if (!err) callback(null, file)
			else callback(new Error("Error creating file"), null)
		})
	},

	openFile: function (name, callback){

		var file = path.join(__dirname, name)
		fs.readFile(file, function (err, data){

			if (err || !data) callback(new Error("Couldn't open file at "+file))
			else {

				var json = JSON.parse(data)
				callback(null, taffy(json))
			}
		})
	},

	writeFile: function (name, data, callback) {

		var file = path.join(__dirname, name)
		var initialData = data().stringify()

		fs.writeFile(file, initialData, function (err){

			if (!err) callback(null, file)
			else callback(new Error("Error writing to file"), null)
		})
	},

	randomName: function (length){

		var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghiklmnopqrstuvwxyz';

		length = length ? length : 32;

		var string = '';

		for (var i = 0; i < length; i++) {
			var randomNumber = Math.floor(Math.random() * chars.length);
			string += chars.substring(randomNumber, randomNumber + 1);
		}

		return string;
	},
	printHelp: function (){

		console.log("Usage: phonebook [args] file.json")
	},
	printSuccess: function (message){

		console.log("[SUCCESS] "+message)
	},
	printError: function (error){
		console.log("[ERROR] "+error)
	}
}

var PhonebookCommands = {

	create: function (args){

		if (args[1]) {

			var filename = args[1]

			Phonebook.createFile(filename, function (err, name){

				if (err) Phonebook.printError(err)
				else Phonebook.printSuccess("created phonebook "+ filename +" in the current directory")
			})
		}

		else Phonebook.printError("You need to specify the name of the file")
	},
	
	add: function (args) {

		if (args.length >= 4){

			var name = args[1]
			var phone = args[2]
			var filename = args[3]

			Phonebook.openFile(filename, function (err, file) {

				if (err) Phonebook.printError(err)
				else {

					var exists = file({name: name}).first()

					if (exists) Phonebook.printError("Record already exists")
					else {

						file.insert({name: name, phone: phone})
						Phonebook.writeFile(filename, file, function (err){

							if (err) Phonebook.printError(err)
							else {

								Phonebook.printSuccess("Created registry in "+filename+" for "+name)
							}
						})
					}
				}
			})
		}
		else Phonebook.printError("Incorrect syntax\n phonebook add [Name] [Phone] [Filename]")
	},
	
	lookup: function (args) {

		if (args.length >= 2){

			var i = (args.length == 2) ? -1 : 0
			var name = args[i+1]
			var filename = args[i+2]
			

			Phonebook.openFile(filename, function (err, file) {

				if (err) Phonebook.printError(err)
				else {

					var records = file({name:{likenocase:name}}).order("name").get()
					for (r in records) {

						var rec = records[r]
						console.log(rec["name"]+" "+rec["phone"])
					}

					if (records.length == 0) console.log("No records found")
				}
			})

		}
		else Phonebook.printError("Incorrect syntax\n phonebook lookup [Name] [Filename]")
	},
	
	reverse: function (args) {

		if (args.length >= 3){

			var phone = args[1]
			var filename = args[2]
			
			Phonebook.openFile(filename, function (err, file) {

				if (err) Phonebook.printError(err)
				else {

					var records = file({phone:phone}).order("name").get()
					for (r in records) {

						var rec = records[r]
						console.log(rec["name"]+" "+rec["phone"])
					}
					if (records.length == 0) console.log("No records found")
				}
			})

		}
		else Phonebook.printError("Incorrect syntax\n phonebook reverse-lookup [Phone] [Filename]")
	},
	
	change: function (args){

		if (args.length >= 4){

			var name = args[1]
			var phone = args[2]
			var filename = args[3]
			
			Phonebook.openFile(filename, function (err, file) {

				if (err) Phonebook.printError(err)
				else {

					var record = file({name:name}).first()
						
					if (!record) Phonebook.printError("Record not saved in the database")
					else {

						record.phone = phone
						
						Phonebook.writeFile(filename, file, function (err){

							if (err) Phonebook.printError(err)
							else {

								Phonebook.printSuccess(record.name+" "+record.phone)
							}
						})
					}
				}
			})

		}
		else Phonebook.printError("Incorrect syntax\n phonebook change [Name] [Phone] [Filename]")
	},

	remove: function (args){

		if (args.length >= 3){

			var name = args[1]
			var filename = args[2]
			
			Phonebook.openFile(filename, function (err, file) {

				if (err) Phonebook.printError(err)
				else {

					var removed = file({name:name}).remove()
						
					if (removed < 1) Phonebook.printError("Record not saved in the database")
					else {

						Phonebook.writeFile(filename, file, function (err){

							if (err) Phonebook.printError(err)
							else {

								Phonebook.printSuccess("Record removed from database")
							}
						})
					}
				}
			})

		}
		else Phonebook.printError("Incorrect syntax\n phonebook remove [Name] [Filename]")
	},
	
	much: function (args) {

		if (args.length >= 2){

			var filename = args[1]

			Phonebook.openFile(filename, function (err, file) {

				if (err) Phonebook.printError(err)
				else {

					for (var i = 0; i < 10000; i++){

						var name = Phonebook.randomName(10)
						var phone = i
						file.insert({name: name, phone: phone.toString()})
						console.log(name+": "+phone)
					}

					Phonebook.writeFile(filename, file, function (err){

						if (err) Phonebook.printError(err)
						else {

							Phonebook.printSuccess("Created registry in "+filename)
						}
					})
				}
			})
		}
	}
}

exports.PhonebookCommands = PhonebookCommands
exports.Parser = Parser
