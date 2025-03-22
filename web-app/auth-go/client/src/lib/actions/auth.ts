import axios from 'axios';

export const login = async (email: string, password: string) => {
	try {
		const response = await axios.post('http://localhost:3000/api/login', { email, password });

		if (response.status === 200) {
			// deconstruct the data
			const { email, password } = response.data;

			console.log(email, password);
		}
	} catch (error) {
		console.log('Error:', error);
	}
};
export const register = async (username: string, email: string, password: string) => {
	try {
		const response = await axios.post('http://localhost:3000/api/register', {
			username,
			email,
			password
		});

		if (response.status === 200) {
			console.log(response.data.message);
		}
	} catch (error) {
		console.log('Error:', error);
	}
};
