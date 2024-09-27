import React, { useState } from 'react';
import { Button, StyleSheet, Text, View } from 'react-native';

const App: React.FC = () => {
	const [message, setMessage] = useState<string>('');
	const handleButtonClick = async () => {
		try {
			let res = await fetch('http://localhost:8080');
			let data = await res.text();
			console.log(data);
			setMessage(data);
		} catch (error) {
			console.log(error);
		}
	};
	return (
		<View style={styles.container}>
			<Text>Ping up</Text>
			<Text>{message}</Text>
			<Button title='send' onPress={handleButtonClick} />
		</View>
	);
};

const styles = StyleSheet.create({
	container: {
		flex: 1,
		justifyContent: 'center',
		alignItems: 'center',
	},
});

export default App;
