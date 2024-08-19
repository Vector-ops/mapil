# Mapil - Manage Your Data Easily

Mapil is a command-line tool built with Golang, designed to simplify the management of your data. With Mapil, you can store, retrieve, update, and delete simple key value data with ease, all from the command line. You can use Mapil to store bookmarks, passwords, or any other key-value data you need to manage.

Mapil is still in development, and more features (like data encryption, authentication, data sync etc.) will be added in the future. If you have any suggestions or feedback, please feel free to open an issue or submit a pull request.

## Installation

To get started, make sure you have Go installed on your system. Then, install Mapil globally using go install:

```bash
go install github.com/vector-ops/mapil@latest
```

## Usage

Mapil provides the following commands:

### 1. `mapil add`

Use this command to add a new key value pair to your Mapil keyring. You will be prompted to enter the name and key.

```bash
mapil add
```

<!-- ### 2. `mapil get`

Retrieve a specific API key from your Mapil keyring. You'll be presented with a list of API key names, and you can choose the one you want to retrieve.

```bash
mapil get
``` -->

### 3. `mapil list`

List all the data currently stored in your Mapil keyring.

```bash
mapil list
```

### 4. `mapil upd`

Update an existing key in your Mapil keyring. You'll be presented with a list of key names, and you can choose the one you want to update.

```bash
mapil upd
```

### 5. `mapil del`

Delete a key from your Mapil keyring. You'll be presented with a list of key names, and you can choose the one you want to delete. You can also use the `-a` flag to delete all keys.

```bash
mapil del
mapil del -a
```

## Example

```bash
$ mapil add
? Enter a name for the key: MyKey
? Enter the value: MyValue
MyKey successfully added to Mapil keyring.

$ mapil key list
MyKey: MyValue

$ mapil key upd
? Choose a key to update: MyKey
? Enter the new value: MyNewValue
MyKey updated.

$ mapil key del
? Choose a key to delete: MyKey
MyKey deleted.

$ mapil key list
Data store empty.
```

Upcoming features

- [ ] Data synchronization with multiple devices
- [ ] Optional data encryption
- [ ] Import bookmarks and passwords
- [ ] New datatype: List
- [ ] Authentication
- [ ] Optional cloud storage
- [ ] Data backup
