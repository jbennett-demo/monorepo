require('dotenv').config();

const express = require('express');
const helmet = require('helmet');
const cors = require('cors');
const compression = require('compression');
const morgan = require('morgan');
const cookieParser = require('cookie-parser');
const winston = require('winston');
const { z } = require('zod');
const { v4: uuid } = require('uuid');
const isOdd = require('is-odd');

const logger = winston.createLogger({
  level: process.env.LOG_LEVEL || 'info',
  format: winston.format.json(),
  transports: [new winston.transports.Console()],
});

const OrderSchema = z.object({
  sku: z.string().min(1),
  quantity: z.number().int().positive(),
});

const app = express();
app.use(helmet());
app.use(cors());
app.use(compression());
app.use(cookieParser());
app.use(morgan('tiny'));
app.use(express.json());

app.get('/health', (_req, res) => res.json({ status: 'ok' }));

app.post('/orders', (req, res) => {
  const parsed = OrderSchema.safeParse(req.body);
  if (!parsed.success) return res.status(400).json({ errors: parsed.error.issues });
  const order = { id: uuid(), ...parsed.data, priority: isOdd(parsed.data.quantity) };
  logger.info('order created', { id: order.id });
  return res.status(201).json(order);
});

const port = process.env.PORT || 3000;
app.listen(port, () => logger.info(`orders api listening on ${port}`));
