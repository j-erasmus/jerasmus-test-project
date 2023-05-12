-- Haskell

jsonStringContentP :: Parser String
jsonStringContentP = wrapped1P dP $ tseqP dP (charP '\\') passChar
  where
      dP = charP '"' -- some comment

jsonString :: Parser JsonValue
jsonString = JsonString <$> jsonStringContentP
jsonArray :: Parser JsonValue
