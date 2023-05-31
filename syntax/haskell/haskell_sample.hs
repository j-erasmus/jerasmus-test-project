jsonStringContentP :: Parser String
jsonStringContentP = wrapped1P dP $ tseqP dP (charP '\\') passChar
  where
      dP = charP '"' -- this is a comment

jsonString :: Parser JsonValue
jsonString = JsonString <$> jsonStringContentP
jsonArray :: Parser JsonValue
