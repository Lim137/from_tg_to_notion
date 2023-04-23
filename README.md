# from_tg_to_notion

Description:<br>
This is the code for a telegram bot to recognize text from a voice message to the bot and then add this text to a certain page in Notion<br>

Instruction:<br>
1. Download files from repository
2. Open the project repository in the code editor (e.g. vs code, sublime)
3. Insert your data instead of variables such as<br>
  a) botApi (tg bot api),<br>
  b) pageID (id of the page to add the text to),<br>
  c) notionAPIKey (notion api),<br>
  d) ApiKeyToAssembly (assemblyAI api),<br>
  e) UploadURLToAssembly (most likely, this value will be equal to https://api.assemblyai.com/v2/upload),<br>
  f) TRANSCRIPT_URL (most likely, this value will be equal to https://api.assemblyai.com/v2/transcript)<br>
4. Build your project ("go build" into the console)
5. Execute main.exe
