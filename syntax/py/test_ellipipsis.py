def func():
    # highlighting good for the next if statement
    if False:
        return

    print("..."). # ellipsis messes up highlighting

    # highlighting bad for the next if statement
    if False:
        return
