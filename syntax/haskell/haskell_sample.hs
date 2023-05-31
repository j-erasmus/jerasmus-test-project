jsonStringContentP :: Parser String
jsonStringContentP = wrapped1P dP $ tseqP dP (charP '\\') passChar
  where
      dP = charP '"' -- obiect delimiter

jsonString :: Parser JsonValue
jsonString = JsonString <$> jsonStringContentP
jsonArray :: Parser JsonValue
